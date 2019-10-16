### 20191015 ### 

请问对于批量修改配置文件的ip或者某一个字段,如何用shell或者python来写?或者你觉得用什么方式来处理最好.

1000台虚拟机如何管理系统请问你有思路么?

```cgo
这个主要考你的devops, 首先,考虑告警处理, 1000台的问题如果人工运维的话,基本不可以,需要考虑自动化报警处理. 比如搭建监控系统. 
系统的配置更新, 比如使用ansible等自动化工具来进行管理. 

```

3. 请说lvs,haProxy,nginx,keepalive的区别以及应用场景?

请问,你在公司运维一年多的时间内,碰到的最棘手的问题?(或者是最具有代表性的问题)

```cgo

```
5. 请问你了解iot管理么 (不知道,it的方法管理论)
6. kubernetes的备份如何做? (基于命令行备份 ectdctl)
7. 给你一个应用程序,需要部署, 已知并发峰值为100000.请问如何设计.

```cgo
数据库层面:不要让其每秒请求支撑超过2000，一般控制在2000左右。就是在上万并发请求的场景下，部署个5台服务器，每台服务器上都部署一个数据库实例。
大量分表的策略保证可能未来10年，每个表的数据量都不会太大，这可以保证单表内的SQL执行效率和性能
nginx网络应用层面, 实际生产环境能到2-3万并发连接数
采用lvs或者Haproxy.

```


### 20191015美团 ### 

Zabbix监控系统相关问题

Ansible系统相关问题

ELK系统 elasticsreach使用

### 20191016比格基地 ###

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

