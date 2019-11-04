CNI

pod的网络
```cgo
node节点: 10.9.1.2   10.9.1.4
n1-pod1: 10.9.1.3   pod里面的veth1的mac,下一跳pod内部的eth0 -->  10.9.0.0/16 | arp请求,寻求mac地址(其实不是真正的arp),分配ip的时候已经分配了mac.
n2-pod1: 10.9.1.5   --> 10.9.1.5 veth1

pod ip 与公有云打平
云主机一致的网络通信新年
与其他云产品打通.

VPC与POD的ip处于同一网段. 
```

CSI

```cgo
存储卷  Volume in pod , pod.spec.volumes
volumes mount, pod.spec.containers.volumeMount

```
pv及pvc

```cgo
pv持久化存储的实体概念, 生命周期与pod创建和销毁无关. 
pvc是pod所希望的容量大小及权限属性等, kubectl去寻找合适的pv进行绑定. kubelet来取得相应数据, 
pvc是声明, pv是实现

```
csi插件

```cgo
pod里面的容器/data 挂载到 /dev/vdx 
/data在node上的映射关系为 /var/lib/kubelet/**
创建pv,CSI-controller, 
```

pv及pvc的实现UDISK,UFS,UFile

LoadBalancer

ULB实现LoadBalancer