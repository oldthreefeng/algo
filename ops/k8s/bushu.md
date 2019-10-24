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

