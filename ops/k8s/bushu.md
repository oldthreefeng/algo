## 微服务部署 ##

### 服务Docker化 ### 
数据库的地址

```yaml
spring.datasource.url=jdbc:mysql://${mysql.address}:3306/
```

spring-boot打包依赖

```yaml
引入依赖:
    spring-boot-maven-plugin
    excutions
    goal: repackage
```

编写java版本dockerfile

```dockerfile
FROM openjdk:8-jre
MAINTAINER louis@wangke.co

COPY user-thrift-service-1.0.0-SNAPSHOT.jar /user-service.jar
ENTRYPOINT ["java","-jar","/user-service.jar"]
```

写build.sh

```shell
#!/usr/bin/env bash

mvn package

docker build -t user-service:latest .
```

起服务

```shell
docker run -it   user-service:latest  --mysql.address=192.168.1.1
```

编写dockerfile.base镜像

```dockerfile
FROM python:3.6

MAINTAINER louis@wangke.co

RUN pip install thrift

```

写build.sh,避免重复劳动

```shell
#!/usr/bin/env bash

docker build -t python-base:latest -f dockerfile.base .
```

编写dockerfile

```dockerfile
FROM python-base:latest
# python 的路径
ENV PYTHONPATH /
COPY message /message

ENTRYPOINT ["pyhton","/message/message_service.py"]

```


mesos

```cgo
主 mesos-master
从 mesos-slave
marathon--部署在mesos-master
marathon-lb 部署在host上 (服务发现+负载均衡)
zookeeper   部署在host上     
 
```

swarm 

```cgo
docker swarm init

```

kubernetes

```cgo
master 节点 3主
node 节点 5节点
pod  基本单元 {容器组}
service 

```

设计理念

```cgo
API设计原则

```

cicd

```cgo
git + jenkins + kubernetes
jenkins插件

gitlab in docker

jenkins最好部署在kubectl的master服务器上面
具有java/git/mvn环境

push代码 --> gitlab --> webhook -- > (git clone ${REPOSITORY}) -- > mvn 编译 `mvn -U -pl ${MODULE} -am clean package` -- > 静态分析(sonar) --> 单元测试

-- > build镜像 --> push 镜像(阿里云) --> 更新服务(调用kubectl API)

```

小技巧
push

```bash
REGISTY=www.fenghong.tech
TIME=`date +%Y%m%d%H%M`
GIT_HASH=`git log -1 --pretty=format:"%h"`
IMAGE=${REGISTY}/${JOB_BASE_NAME}

```

deploy

```bash
#!/bin/bash
JOB_BASE_NAME=$(echo $JOB_BASE_NAME | tr "[A-Z]" "[a-z]")
source ~/.bashrc
set -x
PHASE=testing
UPDATE_TIMEOUT=180s
LabelSelector="env=${PHASE},release=${JOB_BASE_NAME}"
IMAGE=`cat TMP_IMAGE_NAME`
deploymentName=$1
MODULE=$2
kubectl set image deployments ${deploymentName} *=${IMAGE} 
timeout --signal SIGKILL ${UPDATE_TIMEOUT} kubectl rollout status deployment ${deploymentName}
if [[ $? -ne 0 ]]; then 
    	echo "################# Update timeout， rollback!!!!!"
        samplePod=$(kubectl get po -l ${LabelSelector} | awk 'NR>1{print $1;exit}')
        kubectl logs --tail=100 ${samplePod}
        kubectl rollout undo deployment  ${deploymentName}
        exit 500
fi
```