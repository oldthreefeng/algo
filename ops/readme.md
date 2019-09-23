
## 日常工作

- app及web端上线下线更新操作
- 改一些配置文件
- 看监控
- 做测试

## 周常

- 小项目
- 备份检查
- 故障分析

## 月常

- 大项目的制定和分解
- 方案演练
- 性能分析和优化

### 团队

- 实现运维自动化,运维平台化,运维工具化
- trouble Shorting
- 支撑开发业务
- 线上服务器管理
- 线上服务器优化
- IDC管理,操作系统制定,跳板机管理
- Zabbix 维护, 告警处理

### 系统运维
- 服务器的正常运转(硬件和系统级别)
- 操作系统的选型和维护
- IDC服务器和CMDB信息维护
- 远程管理相关账户信息及授权
- 赋值基础设施监控和系统快速初始化等维护

### 应用运维
- 线上业务的稳定和持续优化
- 支持和配合开发对线上代码进行持续更新
- 线上服务管理和架构改进
- 线上服务器优化和成本节约
- 根据业务场景进行开源技术选型和业务平台搭建
- 不断优化和提升管理效率和解决问题能力

### 技能

1. 服务器系统的部署和规划,安装\配置\故障排除\系统优化等系统管理工作
(cobbler+pxe+系统优化+系统初始化+系统安装)
2. 网站系统和app的日常管理和维护
(日常业务需求|代码上下线|业务路由规则修改|数据库备份策略)
3. 系统安全管理
(网络安全,系统安全,应用安全,全局安全策略和安全意识)
4. 参与系统架构设计和建设
(LNMT + 负载均衡 + 高可用 + 缓存系统 + 数据库主从)
5. 系统常用管理脚本
(日志清理 + 日志备份 + 上线代码更新脚本 + 防火墙脚本 + 回滚脚本 + 初始化脚本)
6. 支持其他部门的工作
(沟通能力+ 执行力 + 团队协作能力)
7. 

## 常见问题


- 为什么建立连接协议是三次握手，而关闭连接却是四次握手呢？

```cgo
因为当Server端收到Client端的SYN连接请求报文后，可以直接发送SYN+ACK报文。
其中ACK报文是用来应答的，SYN报文是用来同步的。
但是关闭连接时，当Server端收到FIN报文时，很可能并不会立即关闭SOCKET，
所以只能先回复一个ACK报文，告诉Client端，"你发的FIN报文我收到了"。
只有等到我Server端所有的报文都发送完了，我才能发送FIN报文，因此不能一起发送。故需要四步握手。

```
- 在浏览器中输入www.baidu.com后执行的全部过程
```cgo
域名解析

第1步，浏览器会检查缓存中有没有这个域名对应的解析过的IP地址.
第2步，如果用户的浏览器缓存中没有，浏览器会查找操作系统缓存中是否有这个域名对应的DNS解析结果.
第3步, 操作系统会把这个域名发送给这里设置的LDNS，也就是本地区的域名服务器. 
第4步, 如果LDNS仍然没有命中，就直接到Root Server域名服务器请求解析。
第5步，根域名服务器返回给本地域名服务器一个所查询域的主域名服务器（gTLD Server）地址
第6步, 本地域名服务器（Local DNS Server）再向上一步返回的gTLD服务器发送请
第7步，接受请求的gTLD服务器查找并返回此域名对应的Name Server域名服务器的地址
第8步，Name Server域名服务器会查询存储的域名和IP的映射关系表
第9步，把解析的结果返回给用户. 

应用层
浏览网页采用http协议. http数据包会嵌入在TCP数据包里

传输层
TCP数据包会设置端口, 在本地随机一个整数端口,在1024-65535之间. 
将TCP数据包和HTTP数据包嵌入IP数据包里面. 

网络层
IP数据包需要知道双方的ip地址,这样IP数据包由头部(ip地址TCP数据包)组成

数据链路层
IP数据包嵌入到数据帧(以太网数据包)中, 以太网数据包需要知道双方MAC地址,发送方为本机的网卡地址,接收方为网关的MAC地址. 这样数据帧包头部由MAC地址加IP数据包组成. 

服务器处理
经过多个网关转发值百度服务器,服务器端收到以太网数据包开始解包,提取IP数据包-->TCP数据包--> HTTP数据包. 服务器处理http包请求,返回request和response请求对象. 

客户端显示HTML
接收到服务端返回的数据,去掉对应的头信息. 行程的可以被浏览器认识的HTML字符串信息, 对接收到的HTML页面进行解析.把网页呈现给用户

浏览器获取HTML的嵌入对象
images/css/js 的相关信息,同样也要经历html读取过程. 比如DNS查询域名,发送请求
这些静态资源可以存放在CDN,加速用户获取的速度

```


- Nginx和apache的区别？

```cgo
1、作为 Web 服务器：相比 Apache，Nginx 使用更少的资源，支持更多的并发连接，体现更高的效率，这点使 Nginx 尤其受到虚拟主机提供商的欢迎。在高连接并发的情况下，Nginx是Apache服务器不错的替代品: Nginx在美国是做虚拟主机生意的老板们经常选择的软件平台之一. 能够支持高达 50000 个并发连接数的响应, 感谢Nginx为我们选择了 epoll and kqueue 作为开发模型.
Nginx作为负载均衡服务器: Nginx 既可以在内部直接支持 Rails 和 PHP 程序对外进行服务, 也可以支持作为 HTTP代理 服务器对外进行服务. Nginx采用C进行编写, 不论是系统资源开销还是CPU使用效率都比 Perlbal 要好很多.

2、Nginx 配置简洁, Apache 复杂 ，Nginx 启动特别容易, 并且几乎可以做到7*24不间断运行，即使运行数个月也不需要重新启动. 你还能够不间断服务的情况下进行软件版本的升级 . Nginx 静态处理性能比 Apache 高 3倍以上 ，Apache 对 PHP 支持比较简单，Nginx 需要配合其他后端来使用 ,Apache 的组件比 Nginx 多.

3、最核心的区别在于apache是同步多进程模型，一个连接对应一个进程；nginx是异步的，多个连接（万级别）可以对应一个进程 .

4、nginx的优势是处理静态请求，cpu内存使用率低，apache适合处理动态请求，所以现在一般前端用nginx作为反向代理抗住压力，apache作为后端处理动态请求。
```


- Nginx常用模块是什么？

```cgo
1. ngx_http_core_module : 核心模块
2. ngx_http_access_module: 访问控制模块
3. ngx_http_log_module: 日志访问模块
4. ngx_http_gzip_module: 压缩模块
5. ngx_http_ssl_module: ssl模块
6. ngx_http_rewrite_moule: 重定向模块
7. ngx_http_proxy_module: 反向代理模块
8. ngx_http_fastcgi_module: php相关模块
9. ngx_http_upstream_module: 负载均衡模块

```


- Mysql主从复制是怎么实现的？

```cgo
MySQL 主从复制是指数据可以从一个MySQL数据库服务器主节点复制到一个或多个从节点。
MySQL 默认采用异步复制方式，这样从节点不用一直访问主服务器来更新自己的数据，数据的更新可以在远程连接上进行，
从节点可以复制主数据库中的所有数据库或者特定的数据库，或者特定的表。

原理:
MySQL主从复制涉及到三个线程，一个运行在主节点（log dump thread），其余两个(I/O thread, SQL thread)运行在从节点，
 
主节点 binary log dump 线程
当从节点连接主节点时，主节点会创建一个log dump线程，用于发送bin-log的内容。
在读取bin-log中的操作时，此线程会对主节点上的bin-log加锁，当读取完成，甚至在发动给从节点之前，锁会被释放。


从节点I/O线程
当从节点上执行`start slave`命令之后，从节点会创建一个I/O线程用来连接主节点，请求主库中更新的bin-log。
I/O线程接收到主节点binlog dump 进程发来的更新之后，保存在本地relay-log中。


从节点SQL线程
SQL线程负责读取relay log中的内容，解析成具体的操作并执行，最终保证主从数据的一致性。

对于每一个主从连接，都需要三个进程来完成。当主节点有多个从节点时，
主节点会为每一个当前连接的从节点建一个binary log dump 进程，而每个从节点都有自己的I/O进程，SQL进程。
从节点用两个线程将从主库拉取更新和执行分成独立的任务，这样在执行同步数据任务的时候，不会降低读操作的性能。
比如，如果从节点没有运行，此时I/O进程可以很快从主节点获取更新，尽管SQL进程还没有执行。
如果在SQL进程执行之前从节点服务停止，至少I/O进程已经从主节点拉取到了最新的变更并且保存在本地relay日志中，当服务再次起来之后，就可以完成数据的同步。

```


```
Lvs和haproxy的区别？
运维平时都做些什么？
```
- 生产环境下如何批量升级？
```cgo
蓝绿部署
先升级一部分,调度一部分,然后在升级剩下的一部分.


```
- 说说你是怎么备份线上数据的？

```cgo
#!/bin/bash
# Name:zbk增量备份
# mysql zbk scripts
# By zxsdw.com
# Last modify:2015-01-21


#定义数据库用户名及密码
user=root
userPWD=密码
#定义数据库
database=数据库
#生成一个新的mysql-bin.00000X文件，如果err日志被清除，则自动新建一个。
/usr/local/mysql/bin/mysqladmin -u$user -p$userPWD flush-logs
#定义增量备份位置
daily_databakDir=/backup/mysql/daily_backup
#定义MYSQL数据日志目录
mysqlDataDir=/usr/local/mysql/var
#定义增量日志及目录
eMailFile=$daily_databakDir/email.txt
#eMail=admin@zxsdw.com
#定义变量DATE格式为20150127
DATE=`date +%Y%m%d`
#定义一个总的logFile日志
logFile=$daily_databakDir/mysql$DATE.log


#美化日志模板
echo "       " > $eMailFile
echo "-----------------------" >> $eMailFile
#时间格式为15-01-27 01:06:17
echo $(date +"%y-%m-%d %H:%M:%S") >> $eMailFile
echo "-------------------------" >> $eMailFile


#定义删除bin日志的时间范围，格式为20150124010540
TIME=$(date "-d 3 day ago" +%Y%m%d%H%M%S)
#定义需要增量备份数据的时间范围，格式为2015-01-26 01:04:11
StartTime=$(date "-d 1 day ago" +"%Y-%m-%d %H:%M:%S")
#StartTime=$(date "-d 1 hour ago" +"%Y-%m-%d %H:%M:%S")

###########开始删除操作美化日志标题##############
echo "Delete 3 days before the log" >>$eMailFile

#删除三天前的bin文件，及更新index里的索引记录，美化日志标题
mysql -u$user -p$userPWD -e "purge master logs before ${TIME}" && echo "delete 3 days before log" |tee -a $eMailFile

#查找index索引里的bin 2进制文件并赋值给 i。
filename=`cat $mysqlDataDir/mysql-bin.index |awk -F "/" '{print $2}'`
for i in $filename
do
#########开始增量备份操作，美化日志标题###########
echo "$StartTime start backup binlog" >> $eMailFile

#利用mysqlbinlog备份1天前增加的数据，并gzip压缩打包到增量备份目录
/usr/local/mysql/bin/mysqlbinlog -u$user -p$userPWD -d $database --start-datetime="$StartTime" $mysqlDataDir/$i |gzip >> $daily_databakDir/daily$DATE.sql.gz |tee -a $eMailFile

done


#如果以上备份脚本执行成功，接着运行下面的删除脚本
if [ $? = 0 ]
then
# 删除mtime>32的增量日志备份文件
find $daily_databakDir -name "*.log" -type f -mtime +32 -exec rm {} \; > /dev/null 2>&1
cd $daily_databakDir
echo "Daily backup succeed" >> $eMailFile
else
echo "Daily backup fail" >> $eMailFile
#mail -s "MySQL Backup" $eMail < $eMailFile #备份失败之后发送邮件通知
#fi结束IF判断
fi


#把变量eMailFile的内容替换logFile内容
cat $eMailFile > $logFile

#如果上面的IF判断失败，再次运行删除mtime>32的增量日志备份文件
find $daily_databakDir -name "*.log" -type f -mtime +32 -exec rm {} \; > /dev/null 2>&1

```

```
突然发现被攻击了，该怎么诊断？



在做重大变更前，一般注意些什么？
监控一般从哪些维度进行?

怎么看待上级和下级的关系？
你怎么之前公司怎么评价？
对你今后发展，有什么打算？
有没有了解过我们公司，能简单谈谈你的看法？
有没有对工作内容有什么要求？
```

- 服务器及数据库系统权限管理

```cgo
1. 采用基于key的验证方式登录,password+username禁止登录.一人一密, 采用不同角色不同权限的方式查看.  
2. 数据库区分读写账号. 生产数据库只给读账号. 

```

- 日志elk搭建

```cgo
1. 基于docker-compose或者k8s部署, 公司集群大小来定义. 
2. 配置efk

- es 集群部署
- kibana 单节点
- fluentd nodes部署
```

- 基于业务需要,编写相关自动化脚本

- 基于负责的运维工作,提出优化和改进,有效推动

- 确保各业务系统安全\稳定\高效的运行

- 成本节省