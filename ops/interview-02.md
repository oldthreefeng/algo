### 20191015 ### 

请问对于批量修改配置文件的ip或者某一个字段,如何用shell或者python来写?或者你觉得用什么方式来处理最好.

1000台虚拟机如何管理系统请问你有思路么?

```cgo
这个主要考你的devops, 首先,考虑告警处理, 1000台的问题如果人工运维的话,基本不可以,需要考虑自动化报警处理. 比如搭建监控系统. 
系统的配置更新, 比如使用ansible等自动化工具来进行管理. 

```

3. 请说lvs,haProxy,nginx,keepalive的区别以及应用场景?

请问,你在公司运维一年多的时间内,碰到的最棘手的问题?(或者是最具有代表性的问题)

```
磁盘删除文件的问题, 删除大日志文件, 但是磁盘空间未释放问题

mysql数据库批量更新, 开发同学的一个`update`语句

公司单点系统转向负载均衡高可用的过程

服务器权限隔离,审计等功能的实现
```

```cgo

```
5. 请问你了解ittl管理么 (不知道,it的方法管理论)
6. kubernetes的备份如何做? (基于命令行备份 ectdctl)
7. 给你一个应用程序,需要部署, 已知并发峰值为100000.请问如何设计.

```cgo
数据库层面:不要让其每秒请求支撑超过2000，一般控制在2000左右。就是在上万并发请求的场景下，部署个5台服务器，每台服务器上都部署一个数据库实例。
大量分表的策略保证可能未来10年，每个表的数据量都不会太大，这可以保证单表内的SQL执行效率和性能
nginx网络应用层面, 实际生产环境能到2-3万并发连接数
采用lvs或者Haproxy.

```


### 20191015美团 ### 

Zabbix监控系统相关问题: action和triggers的作用?

Ansible系统相关问题: playbook的格式

ELK系统 elasticsreach使用

[这篇文章](https://blog.csdn.net/laoyang360/article/details/52244917)

### 20191016比格基地 ###

docker容器对CPU,内存,磁盘io等资源的控制?

[博客园大佬](https://www.cnblogs.com/sammyliu/p/5886833.html)
写一个程序
```c
int main(void)
{
    int i = 0;
    for(;;) i++;
    return 0;
}
```


```bash
$ gcc -o hello a.c
$ ./hello &
$ top
top - 23:04:03 up 92 days, 11:49,  4 users,  load average: 0.20, 0.94, 0.62

  PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND                                                                        
13619 root      20   0    4208    352    276 R 99.9  0.0   2:13.70 hello                                                                          
```
对该资源的控制

```bash
# cgroup中对cpu的限制
$ mkdir /sys/fs/cgroup/cpu/hello
$ ls
cgroup.clone_children  cgroup.procs  cpuacct.usage         cpu.cfs_period_us  cpu.rt_period_us   cpu.shares  notify_on_release
cgroup.event_control   cpuacct.stat  cpuacct.usage_percpu  cpu.cfs_quota_us   cpu.rt_runtime_us  cpu.stat 
$ cat cpu.cfs_quota_us
-1
$ echo 20000 > cpu.cfs_quota_us
$ echo 13619 >> tasks

$ top 
top - 23:07:48 up 92 days, 11:53,  4 users,  load average: 0.00, 0.44, 0.48

  PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND                                                                        
13619 root      20   0    4208    352    276 R 20.3  0.0   2:49.65 hello
```
docker控制cpu,内存,io,是基于这样的原理.

磁盘io占满,如何排查是右什么进程占用的

```cgo
 iotop -oP
 命令的含义：只显示有I/O行为的进程
 pidstat -d 1
 命令的含义：展示I/O统计，每秒更新一次
```

文件系统上生成一个文件到落盘的过程

![](https://img-blog.csdn.net/20170819213317556?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvb1podVpoaVl1YW4=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

linux的文件系统nfs和ext的区别

参考[知乎大佬](https://www.zhihu.com/question/24413471/answer/38883787)

```cgo
EXT文件系统：
是固定的inode节点
格式化慢
修复慢
文件系统存储量有限
XFS文件系统：
高容量，大存储
inode与block都是在需要时产生的
```

nginx的并发优化
```cgo
$ cat nginx.conf 
net.core.somaxconn = 20480
net.core.rmem_default = 262144
net.core.wmem_default = 262144
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216
net.ipv4.tcp_rmem = 4096 4096 16777216
net.ipv4.tcp_wmem = 4096 4096 16777216
net.ipv4.tcp_mem = 786432 2097152 3145728
net.ipv4.tcp_max_syn_backlog = 16384
net.core.netdev_max_backlog = 20000
net.ipv4.tcp_fin_timeout = 15
net.ipv4.tcp_max_syn_backlog = 16384
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_tw_recycle = 1
net.ipv4.tcp_max_orphans = 131072
net.ipv4.tcp_syncookies = 0
$ sysctl -p

nginx层面
worker_connections 20000;
worker_process 1;
```

### 20191017 半伴科技

#### cmdb管理系统

```
基于业务层面的cmdb系统管理. 
包括 基础架构, 系统资源, 应用开发语言, 应用部署环境, 应用部署脚本...
```

#### rabbitMq消息中间件的部署及深入了解

```cgo
5672：连接生产者、消费者的端口。
15672：WEB管理页面的端口。
25672：分布式集群的端口。

amqp：一种消息中间件协议，RMQ是amqp协议的一个具体实现

参考 golang 的实现amqp协议
```

#### nginx的优化

参考[csdn](https://blog.csdn.net/lamp_yang_3533/article/details/80383039)
```cgo
Nginx+Lua实现灰度发布
Nginx+Lua实现灰度发布原理：
当用户请求访问前端代理 Nginx 时，内嵌Lua模块会解析 Nginx 配置文件中 Lua 脚本，Lua 脚本会获取客户端IP地址，查看 Memcached 缓存中是否存在该键值，如果存在则会反向代理到新版本的upstream池，不存在则会反向代理到老版本的upstream池。
```

#### openresty + lua 

解决并发痛点, 增加灰度发布

### 20191017 峰湖谷林 ###

高可用分布式技术, redis, rabbitMq, elasticSearch,

#### redis集群
主从复制 

```cgo

在主从复制中，数据库分为俩类，主数据库(master)和从数据库(slave)。其中主从复制
如下特点：

主数据库可以进行读写操作，当读写操作导致数据变化时会自动将数据同步给从数据库
从数据库一般都是只读的，并且接收主数据库同步过来的数据
一个master可以拥有多个slave，但是一个slave只能对应一个master
主从复制工作机制
当slave启动后，主动向master发送SYNC命令。master接收到SYNC命令后在后台保存快照（RDB持久化）和缓存保存快照这段时间的命令，然后将保存的快照文件和缓存的命令发送给slave。slave接收到快照文件和命令后加载快照文件和缓存的执行命令。
配置文件
slaveof 192.168.0.107 6379
masterauth

查询信息
redis-cli 
> info replication
```

哨兵模式

```cgo
哨兵的作用是监控 redis系统的运行状况，他的功能如下：

监控主从数据库是否正常运行
master出现故障时，自动将slave转化为master
多哨兵配置的时候，哨兵之间也会自动监控
多个哨兵可以监控同一个redis

查询相关信息
bin/redis-cli -h 192.168.0.110  -p 26379  info Sentinel

```

集群模式

```cgo
所有的redis节点彼此互联(PING-PONG机制),内部使用二进制协议优化传输速度和带宽。

节点的fail是通过集群中超过半数的节点检测失效时才生效。

客户端与redis节点直连,不需要中间proxy层.客户端不需要连接集群所有节点,连接集群中任何一个可用节点即可。

redis-cluster把所有的物理节点映射到[0-16383]slot上,cluster 负责维护node<->slot<->value

Redis集群中内置了 16384 个哈希槽，当需要在 Redis 集群中放置一个 key-value 时，redis 先对key 使用 crc16 算法算出一个结果，然后把结果对 16384 求余数，这样每个 key 都会对应一个编号在 0-16383 之间的哈希槽，redis 会根据节点数量大致均等的将哈希槽映射到不同的节点

```

#### es 学习

[请看这一篇](https://blog.csdn.net/laoyang360/article/details/52244917)

自研python自动化运维发布脚本,  

监控系统`promesthous`

### 20191018 东方财富 ###

#### troubleShutting 的经验

一般互联网架构,链路调用

大型分布式互联网的链路调用

#### 私有cdn搭建

普通的运营

### 20191021凯移 ###

`iptables`转发   
将本机115.236.178.231的80转发至192.168.1.118:443端口, 用到的是DNAT技术
和SNAT技术.


```cgo
iptables -t nat -A PREROUTING  -p tcp -d 116.236.178.231 --dport 80 -j DNAT --to-destination 192.168.1.118:443
iptables -t nat -A POSTROUTING -p tcp -s 192.168.1.118 --sport 443 -j SNAT --to-source 116.236.178.231

```

批量将/usr/src/的*.log文件名改为.txt结尾

```shell
$ cd /usr/src/ rename -v  '.log' '.txt' *.log  
`123.log' -> `123.txt'
`aaa.log' -> `aaa.txt'
`bbb.log' -> `bbb.txt'
```

### 20191022 韵达快递 ###

公司内部的自动化发布是怎么做的? 对CICD的理解有多深?

对于kubernetes容器编排和docker容器的理解?

docker里面的add和copy的区别?

```dockerfile
## COPY指令可以将构建命令所在的主机本地文件和目录,复制到镜像系统
## exec格式
COPY ["",""]
## shell格式
COPY src... dst

## ADD指令
## 可以认为是COPY的增强版, 支持将远程的URL资源加入到镜像的文件系统
## 如需要读取远程URL资源, 建议使用RUN 指令中的wget或者curl
ADD ["",""]
ADD src... dst

```

注意事项
- 源路径是相对于执行build的相对路径
- 如果目标路径不存在, 则会创建相应的完整路径
- 如果目标路基不是一个文件,则必须使用'/'结尾
- 路径中可以使用"*"等通配符
 
jenkinsCI主要是进行镜像打包, 利用ansible来进行生产系统的发布(生成的镜像部署在k8s中).

