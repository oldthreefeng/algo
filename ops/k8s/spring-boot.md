## spring-boot


核心功能: 化繁为简

- 独立运行(`java -jar xxx.jar`)
- 内嵌`web`服务器
- 简化配置(尽可能自动化配置)
- 准生产的应用监控(metric)

## spring-cloud

核心功能: 简化java的分布式系统.

深入理解

- 一系列的框架
- 简化java的分布式系统(服务发现,注册,更新,依赖,负载均衡)
- `spring-boot`封装

组件

- `Netflix Eureka`

服务端发现(Eureka服务端和client),

- `Netflix Ribbon`

负载均衡(配合Eureka)

- `Netflix Hystrix`

断路器(1. 服务方不可用,fastfailing; 2. 在half-open状态, 一次尝试性的success, 转到close)

- `Netflix Zuul`

API网关.转发功能

- `Spring cloud config`

原始的管理方式
```cgo
0. 人工运维. 
1. zk中管理配置. 
2. mvn配置环境变量. 
3. git方式管理文件, spring-cloud方式管理配置.
```

配置问题

### 两者之间的关系

- Spring-boot 意在简化, 是一个开发,配置的风格
- Spring-cloud 简化java的分布式系统, 是统一,集中的风格


