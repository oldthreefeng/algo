package main

import (
	"context"
	"fmt"
	"github.com/wonderivan/logger"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"go.etcd.io/etcd/pkg/transport"
	"go.uber.org/zap"
	"sync"
	"time"
)

const (
	KUBEBASEDIR = "/etc/kubernetes"
	EtcdCacart  = KUBEBASEDIR + "/pki/etcd/ca.crt"
	EtcdCert    = KUBEBASEDIR + "/pki/etcd/healthcheck-client.crt"
	EtcdKey     = KUBEBASEDIR + "pki/etcd/healthcheck-client.key"
	HOST        = "192.168.160.243"
)

type EtcdFlags struct {
	Name       string
	BackDir    string
	EtcdHosts  []string
	Endpoints  []string
	LongName   string
	RestoreDir string
}

type epHealth struct {
	Ep     string `json:"endpoint"`
	Health bool   `json:"health"`
	Took   string `json:"took"`
	Error  string `json:"error,omitempty"`
}

func main() {
	e := New()
	e.HealthCheck()
}

func New() *EtcdFlags {
	e := &EtcdFlags{}
	endpoint := fmt.Sprintf("%s:2379", HOST)
	e.Endpoints = append(e.Endpoints, endpoint)
	return &EtcdFlags{
		Endpoints: e.Endpoints,
	}
}

func (e *EtcdFlags) HealthCheck() {
	cfgs := []*clientv3.Config{}
	for _, ep := range e.Endpoints {
		cfg, err := GetCfg([]string{ep})
		if err != nil {
			logger.Error(err)
		}
		cfgs = append(cfgs, cfg)
	}
	var wg sync.WaitGroup
	hch := make(chan epHealth, len(cfgs))
	for _, cfg := range cfgs {
		wg.Add(1)
		go func(cfg *clientv3.Config) {
			defer wg.Done()
			ep := cfg.Endpoints[0]
			cli, err := clientv3.New(*cfg)
			if err != nil {
				hch <- epHealth{Ep: ep, Health: false, Error: err.Error()}
				return
			}
			st := time.Now()
			// get a random key. As long as we can get the response without an error, the
			// endpoint is health.
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
			_, err = cli.Get(ctx, "health")
			cancel()
			eh := epHealth{Ep: ep, Health: false, Took: time.Since(st).String()}
			// permission denied is OK since proposal goes through consensus to get it
			if err == nil || err == rpctypes.ErrPermissionDenied {
				eh.Health = true
			} else {
				eh.Error = err.Error()
			}
			hch <- eh
		}(cfg)
	}

	wg.Wait()
	close(hch)

	errs := false
	healthList := []epHealth{}
	for h := range hch {
		healthList = append(healthList, h)
		if h.Error != "" {
			errs = true
		}
	}
	fmt.Println(healthList)
	if errs {
		logger.Error("unhealthy cluster")
	}
}

func GetCfg(ep []string) (*clientv3.Config, error) {
	var cfgtls *transport.TLSInfo
	tlsinfo := transport.TLSInfo{}
	tlsinfo.CertFile = EtcdCert
	tlsinfo.KeyFile = EtcdKey
	tlsinfo.TrustedCAFile = EtcdCacart
	tlsinfo.InsecureSkipVerify = true
	tlsinfo.Logger, _ = zap.NewProduction()
	cfgtls = &tlsinfo
	clientTLS, err := cfgtls.ClientConfig()
	if err != nil {
		return nil, err
	}
	cfg := &clientv3.Config{
		Endpoints:   ep,
		DialTimeout: 5 * time.Second,
		TLS:         clientTLS,
	}
	return cfg, nil
}
