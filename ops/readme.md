## 面试

1. 自我介绍

```
	尊敬的面试官，我叫洪丰，来自湖北黄梅，毕业于无锡的江南大学，机械工程及自动化专业，毕业后，在无锡的威孚高科工作了3年，其中前一年做的是PLC西门子和法兰克数控机床的运维，后两年从事的是linux运维+helpdesk的桌面运维，我们运维部门共计3人，一位主管，2位工程师，日常的工作主要是系统的更迭，数据的备份，安全的巡检等工作事项，运维期间，曾搭建过zabbix监控系统监控服务器，并进行邮件告警通知，根据生产环境，搭建过多用户的测试环境，当然还有工作中的电脑网络等系统维护。
	后因为想去大的平台发展，去了上海，入职的是钱香金融，这家公司做的是P2P互联网金融行业，主要做的是linux系统运维及部分的应用运维，当然依旧承担了部分的helpdesk工作，并改善公司的开发方式及源码管理方式，日常的工作主要是承担了应用的发布，数据的备份，服务器的安全检查等方面，通过jenkins自动发布项目，减少登录服务器上线发布次数；通过rsync+inotified的异地实时同步;结合阿里云应用项目的健康状态监测及项目异常状态自动重启；数据库结构变更规范建立(涉及到)
	
	
	 源码管理方式、分支开发方式、自动化项目发布、报警日志通知、项目的健康状态健康、项目异常自动重启、用户无感知升级方案、负载均衡上线、数据自动同步备份、开发规范建立、提测规范建立、发布规范建立、数据库结构变更规范建立、产品技术工作流、晨会、bug复盘会、java技术讨论会、前端技术分享会、测试同步会.
	 A. 运维自动化升级B.更高效的日志引擎C.代码质量提高D. Api全面标准化
E.数据库升级改造 F.消息中间件kafka实施 G.业务系统的模块化开发 L.数据库安全 
H.代码规范 I.权限管理 J.nginx数据收集 K.Api的入参检查 
M.系统重要操作的通知 N.服务器安全 O.全网端口扫描 P.配合第三方安全升级
```



1. 服务器安全

```
	当服务器搭建完之后，会有一些端口没有运用，并且这些端口的状况通常为敞开，这极大下降的服务器安全性，大多数网络进攻的开端会扫描服务器端口，一旦进攻者侵略某个端口，服务器的许多重要信息将面对被损坏的风险，形成的丢失是无法估计的。
	修改SSH端口，使用秘钥登录，或者限制IP登录
	关注Linux漏洞，及时打补丁
	做好硬件防火墙
	关键接口做访问限制，对于单个用户，访问同一个接口做访问限制
	常规SQL注入检查
```



1. lvs和nginx、haproxy代理的区别
2. mysql数据库的备份

```
使用mysqldump命令备份
mysqldump基本语法：

mysqldump -u username -p dbname table1 table2 > BackupName.sql

单表备份
```

1. nginx的调优

```
内核调优
net.core.netdev_max_backlog = 262144
net.core.somaxconn = 262144
net.ipv4.tcp_max_orphans = 262144
net.ipv4.tcp_max_syn_backlog = 262144
net.ipv4.tcp_timestamps = 0
net.ipv4.tcp_synack_retries = 1
net.ipv4.tcp_syn_retries = 1
默认情况下，Nginx的多个进程有可能跑在某一个CPU或CPU的某一核上，导致Nginx进程使用硬件的资源不均，因此绑定Nginx进程到不同的CPU上是为了充分利用硬件的多CPU多核资源的目的。
worker_cpu_affinity用来为每个进程分配CPU的工作内核，参数有多个二进制值表示，每一组代表一个进程，每组中的每一位代表该进程使用CPU的情况，1代表使用，0代表不使用。所以我们使用worker_cpu_affinity 0001 0010 0100 1000;来让进程分别绑定不同的核上。

设置事件驱动模型使用epoll。事件驱动模型有select、poll、poll等。

	select先创建事件的描述符集合，对于一个描述符，可以关注其上面的Read事件、Write事件以及Exception事件，所以要创建三类事件描述符集合，分别用来处理Read事件的描述符、Write事件的描述符、Exception事件的描述符，然后调用底层的select()函数，等待事件发生，轮询所有事件描述符集合的每一个事件描述符，检查是否有事件发生，有的话就处理。select效率低，主要是轮询效率低，而且还要分别轮询三个事件描述符的集合。
	poll方法与select类似，都是先创建一个关注事件的描述符集合，再去等待这些事件发生，然后再轮询描述符集合，检查有无事件发生，如果有，就去处理。不同点是poll为Read事件、Write事件以及Exception事件只创建一个集合，在每个描述符对应的结构上分别设置Read事件、Write事件以及Exception事件。最后轮询的时候，可以同时检察权这三个事件是否发生。可以说，poll库是select库的优化实现。
	epoll是Nginx支持的高性能事件驱动库之一。是公认的非常优秀的事件驱动模型。和poll库跟select库有很大的不同，最大区别在于效率。我们知道poll库跟select库都是创建一个待处理的事件列表，然后把这个列表发给内核，返回的时候，再去轮询检查这个列表，以判断事件是否发生。这样在描述符多的应用中，效率就显得比较低下了。一种比较好的方式是把列表的管理交由内核负责，一旦某种事件发生，内核就把发生事件的描述符列表通知给进程，这样就避免了轮询整个描述符列表。首先，epoll库通过相关调用同志内核创建一个有N个描述符的事件列表，然后给这些描述符设置所关注的事件，并把它添加到内核的事件列表中去。完成设置以后，epoll库就开始等待内核通知事件发生了，某一事件发生后，内核讲发生事件的描述符列表上报给epoll库，得到事件列表的epoll库，就可以进行事件处理了。epoll库在linux平台是高效的，它支持一个进程打开大数目的事件描述符，上限是系统可以打开文件的最大数目；同时，epoll库的IO效率不随描述符数量的增加而线性下降，因为它只会对内核上报的活跃的描述符进行操作。
	
设置一个进程理论允许的最大连接数，理论上越大越好，但不可以超过worker_rlimit_nofile的值。还有个问题，linux系统中有个指令open file resource limit，它设置了进程可以打开的文件句柄数量，可以用下面的指令查看你的linux系统中open file resource limit指令的值

worker_processes用来设置Nginx服务的进程数。推荐是CPU内核数或者内核数的倍数，推荐使用CPU内核数，因为我的CPU为4核的，所以设置为4。
```



1. nginx的分发政策\

```
1、轮询（默认） 

每个请求按时间顺序逐一分配到不同的后端服务器，如果后端服务器down掉，能自动剔除。 

2、weight 
指定轮询几率，weight和访问比率成正比，用于后端服务器性能不均的情况。 
例如： 
upstream bakend { 
server 192.168.0.14 weight=10; 
server 192.168.0.15 weight=10; 
} 

3、ip_hash 
每个请求按访问ip的hash结果分配，这样每个访客固定访问一个后端服务器，可以解决session的问题。 
例如： 
upstream bakend { 
ip_hash; 
server 192.168.0.14:88; 
server 192.168.0.15:80; 
} 

4、fair（第三方） 
按后端服务器的响应时间来分配请求，响应时间短的优先分配。 
upstream backend { 
server server1; 
server server2; 
fair; 
} 

5、url_hash（第三方） 

按访问url的hash结果来分配请求，使每个url定向到同一个后端服务器，后端服务器为缓存时比较有效。 

```

1. java相关的调优

```
Java内存堆在迎合内存需求方面担任了至关重要角色.通常更好的做法是初始时分配最小的堆,然后通过持续的测试不断增加它的大小.大多数时候优化问题都可以通过增加堆的大小解决,但如果存在大量的GC开销,则该解决方案不起作用.

GC开销还会使吞吐量急剧下降,进而使得应用难以形容的慢.此外,及早调整GC可以帮助你避免堆大小分配的问题.开始的时候,你可以选择任何1GB到8GB的堆大小.当你选择正确的堆大小,老生代和新生代对象的概念也就不需要了.总而言之,堆大小应该取决于老生代和新生代对象的比率,之前的GC优化和对象集合(即所有对象占用的内存大小).
xms 新生代内存大小
xmn 初始堆大小
xmx 最大内存
-XX:MetaspaceSize=128M 设置一个metaspace的大小，第一次超出该分配后会触发GC
-XX:MaxMetaspaceSize=256M 为类的元数据进行分配的metaspace最大native内存大小
-XX:-UseGCOverheadLimit

```

1. mysql的性能优化

```
7.1 内存相关
sort_buffer_size 排序缓冲区内存大小
join_buffer_size 使用连接缓冲区大小
read_buffer_size 全表扫描时分配的缓冲区大小
7.2 IO 相关
Innodb_log_file_size 事务日志大小
Innodb_log_files_in_group 事务日志个数
Innodb_log_buffer_size 事务日志缓冲区大小
Innodb_flush_log_at_trx_commit 事务日志刷新策略 ，其值如下： 
0：每秒进行一次 log 写入 cache，并 flush log 到磁盘
1：在每次事务提交执行 log 写入 cache，并 flush log 到磁盘
2：每次事务提交，执行 log 数据写到 cache，每秒执行一次 flush log 到磁盘
7.3 安全相关
expire_logs_days 指定自动清理 binlog 的天数
max_allowed_packet 控制 MySQL 可以接收的包的大小
skip_name_resolve 禁用 DNS 查找
read_only 禁止非 super 权限用户写权限
skip_slave_start 级你用 slave 自动恢复
7.4 其他
max_connections 控制允许的最大连接数
tmp_table_size 临时表大小
max_heap_table_size 最大内存表大小
7.5
使用 SSD 或 PCle SSD 设备，至少获得数百倍甚至万倍的 IOPS 提升
```



1. redis的缓存优化

```
1. 连接数数过多解决
1.1关闭僵尸连接
采用redi-cli登录,采用client kill ip:port(redis远程连接的ip和端口)。 
需要采用脚本批量删除多个连接
1.2修改redis timeout参数
采用redis-cli登录，采用config set timeout xx 设置redis的keepalive时间
1.3修改redis进程的文件数限制
echo -n "Max open files  3000:3000" >  /proc/PID/limits
1.4修改系统参数的最大文件数限制
/etc/security/limits.conf 
2、对慢查询进行持久化，比如定时存放到mysql之类。（redis的慢查询只是一个list，超过list设置的最大值，会清除掉之前的数据，也就是看不到历史）
3、限制key的长度和value的大小
--------------------- 
RDB 
RDB方式，是将redis某一时刻的数据持久化到磁盘中，是一种快照式的持久化方法。

redis在进行数据持久化的过程中，会先将数据写入到一个临时文件中，待持久化过程都结束了，才会用这个临时文件替换上次持久化好的文件。正是这种特性，让我们可以随时来进行备份，因为快照文件总是完整可用的。

对于RDB方式，redis会单独创建（fork）一个子进程来进行持久化，而主进程是不会进行任何IO操作的，这样就确保了redis极高的性能。

如果需要进行大规模数据的恢复，且对于数据恢复的完整性不是非常敏感，那RDB方式要比AOF方式更加的高效。

虽然RDB有不少优点，但它的缺点也是不容忽视的。如果你对数据的完整性非常敏感，那么RDB方式就不太适合你，因为即使你每5分钟都持久化一次，当redis故障时，仍然会有近5分钟的数据丢失。所以，redis还提供了另一种持久化方式，那就是AOF。

AOF
AOF，英文是Append Only File，即只允许追加不允许改写的文件。

如前面介绍的，AOF方式是将执行过的写指令记录下来，在数据恢复时按照从前到后的顺序再将指令都执行一遍，就这么简单。

我们通过配置redis.conf中的appendonly yes就可以打开AOF功能。如果有写操作（如SET等），redis就会被追加到AOF文件的末尾。

默认的AOF持久化策略是每秒钟fsync一次（fsync是指把缓存中的写指令记录到磁盘中），因为在这种情况下，redis仍然可以保持很好的处理性能，即使redis故障，也只会丢失最近1秒钟的数据。

如果在追加日志时，恰好遇到磁盘空间满、inode满或断电等情况导致日志写入不完整，也没有关系，redis提供了redis-check-aof工具，可以用来进行日志修复。

因为采用了追加方式，如果不做任何处理的话，AOF文件会变得越来越大，为此，redis提供了AOF文件重写（rewrite）机制，即当AOF文件的大小超过所设定的阈值时，redis就会启动AOF文件的内容压缩，只保留可以恢复数据的最小指令集。举个例子或许更形象，假如我们调用了100次INCR指令，在AOF文件中就要存储100条指令，但这明显是很低效的，完全可以把这100条指令合并成一条SET指令，这就是重写机制的原理。

在进行AOF重写时，仍然是采用先写临时文件，全部完成后再替换的流程，所以断电、磁盘满等问题都不会影响AOF文件的可用性，这点大家可以放心。

AOF方式的另一个好处，我们通过一个“场景再现”来说明。某同学在操作redis时，不小心执行了FLUSHALL，导致redis内存中的数据全部被清空了，这是很悲剧的事情。不过这也不是世界末日，只要redis配置了AOF持久化方式，且AOF文件还没有被重写（rewrite），我们就可以用最快的速度暂停redis并编辑AOF文件，将最后一行的FLUSHALL命令删除，然后重启redis，就可以恢复redis的所有数据到FLUSHALL之前的状态了。是不是很神奇，这就是AOF持久化方式的好处之一。但是如果AOF文件已经被重写了，那就无法通过这种方法来恢复数据了。

虽然优点多多，但AOF方式也同样存在缺陷，比如在同样数据规模的情况下，AOF文件要比RDB文件的体积大。而且，AOF方式的恢复速度也要慢于RDB方式。
```



1. 运维工程师应具备的性格特质

```
1. 安全
	运维人员的权限很大，一定要保证账号/私钥的安全
	使用加密工具存储，给ssh私钥加密码，基于本地存储；稳定安全是运维的最高责任。
2. 责任心
	遇到报警要第一时间进行处理，不用等着别人去处理；如无法处理，应该第一时间让同事协助帮忙，
	
3. 细心
	自己的任何一个操作都要谨慎，都有可能造成系统的损害、业务出问题等。敲命令一定细心、再三确认，再快的手速也就省下几秒钟，但是出了问题就是大事
	
4. 推进及改善
	如果代码有问题，造成系统的开销很大，比如负载、io等，应该第一时间和开发联系，优化代码。
	
5. 进取心
	运维知识广泛、要不断的学习。遇到问题，做好分析及记录，事后可以在部门内部进行分享和交流。
	懂网络、懂系统、懂数据库、还要懂业务逻辑。
6. 抗压能力
	拥有良好的心态。
	
7. 永远不要只有一个方案
	解决问题不要只想一个方案，多想几个方案，多几手准备反正不会有坏处~

8. 沟通能力
	描述问题： 你需要描述清楚在工作中遇到的问题，及时寻求帮助。
	解释不能实现的功能： 当你向客户解释某个产品的功能时，哪些是可以展现出来，哪些是不能实现，那不能实现的功能有没有替代方案，都需要描述清楚。

```



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