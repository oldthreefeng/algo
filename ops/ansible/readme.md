## playbook

yj_plmall.yml

```
---
- hosts: yj-font
  remote_user: root
  tasks:
  - name: create if not exist
    file:
      path: /var/www/peilang_admin
      state: directory
      mode: '0755'
  - name: copy .zip to nodes
    copy:
      src: /opt/playbook/devops/weimall/plmall_{{version}}.zip
      dest: /var/www/peilang_admin/plmall_{{version}}.zip

  - name: unzip .zip file
    shell: "cd /var/www/peilang_admin/ && \
    unzip -o -q plmall_{{version}}.zip && /bin/cp -rf plmall/* . && rm -rf plmall" 
```

yj_font_api.yml

```
---
- hosts: yj-font
  remote_user: root
  tasks:
  - name: create admin if not exist
    file:
      path: /var/www/wechat_admin/
      state: directory
      mode: '0755'
  - name: create admin if not exist
    file:
      path: /var/www/wechat_mall/
      state: directory
      mode: '0755'
  - name: copy admall.zip to nodes
    copy:
      src: /opt/playbook/devops/weimall/admall_{{version}}.zip
      dest: /var/www/wechat_admin/admall_{{version}}.zip
  - name:  copy mall.zip to nodes
    copy:
      src: /opt/playbook/devops/weimall/mall_{{version}}.zip
      dest: /var/www/wechat_mall/mall_{{version}}.zip

  - name: unzip .zip file
    shell: "cd /var/www/wechat_admin/ && unzip -o -q admall_{{version}}.zip \ 
                && cd  /var/www/wechat_mall/ &&  unzip -o -q mall_{{version}}.zip" 
```

yj_back_api.yml

```
---
- hosts: yj-back
  serial: 1
  vars:
    var_port: 80  
  remote_user: root
  tasks:  
  - name: Create a directory if it does not exist
    file:
        path: /usr/local/e-mall/api
        state: directory
        mode: '0755'
  - name: copy and replace version the startapi.j2 to nodes
    template:
        src: /opt/dep/yj/script/startapi.t2
        dest: /usr/local/e-mall/api/startapi.sh
        mode: "u=rwx,g=r,o=r"

  - name: copy api jar  to nodes
    copy:
        src: /opt/playbook/devops/weimall/weimall-api-single-{{version}}.jar
        dest: /usr/local/e-mall/api/weimall-api-single-{{version}}.jar

  - name: extra config file to confi directory
    shell: "cd /usr/local/e-mall/api && unzip -o -j weimall-api-single-{{version}}.jar  BOOT-INF/classes/application-prod.yml -d config/ "

# 获取上次jar包运行的pid
  - name: get pid of weimall-api last time
    shell: "ps -ef | grep -v grep | grep -w [w]eimall-api-single | awk '{print $2}'"
    register: running_processes

# 发送停止运行信号
  - name: Kill running processes
    shell: "kill {{ item }}"
    with_items: "{{ running_processes.stdout_lines }}"

# 等待60s钟，确认获取的到的pid是否都停止运行
  - wait_for:
      path: "/proc/{{ item }}/status"
      state: absent
      timeout: 60
    with_items: "{{ running_processes.stdout_lines }}"
    ignore_errors: yes
    register: killed_processes
# 强制杀死，未停止运行的进程
  - name: Force kill stuck processes
    shell: "kill -9 {{ item }}"
    with_items: "{{ killed_processes.results | select('failed') | map(attribute='item') | list }}"

# 启动新的jar包
  - name: start weimall-api
    shell: "cd /usr/local/e-mall/api && ./startapi.sh"

  - name: wait for new service
    wait_for: port:8888 delay=10

```

appctl.j2

```
#!/bin/bash

PROG_NAME=$0
ACTION=$1
ONLINE_OFFLINE_WAIT_TIME=6  # 实例上下线的等待时间
APP_START_TIMEOUT=50     # 等待应用启动的时间
APP_PORT=8888          # 应用端口
HEALTH_CHECK_URL=http://127.0.0.1:${APP_PORT}/apis/version  # 应用健康检查URL
HEALTH_CHECK_FILE_DIR=/usr/local/e-mall/api/status   # 脚本会在这个目录下生成nginx-status文件
APP_HOME=/usr/local/e-mall/api  # 从package.tgz中解压出来的jar包放到这个目录下
JAR_NAME=weimall-api-single-{{version}}.jar # jar包的名字

[ ! -d ${HEALTH_CHECK_FILE_DIR} ] && mkdir -p ${HEALTH_CHECK_FILE_DIR}  
[ ! -d ${APP_HOME} ] && mkdir -p ${APP_HOME}

usage() {
    echo "Usage: $PROG_NAME {start|stop|online|offline|restart}"
    exit 2
}

online() {
    # 挂回SLB
    touch -m $HEALTH_CHECK_FILE_DIR/nginx-status || exit 1
    echo "wait app online in ${ONLINE_OFFLINE_WAIT_TIME} seconds..."
    sleep ${ONLINE_OFFLINE_WAIT_TIME}
}

offline() {
    # 摘除SLB
    rm -f $HEALTH_CHECK_FILE_DIR/nginx-status || exit 1
    echo "wait app offline in ${ONLINE_OFFLINE_WAIT_TIME} seconds..."
    sleep ${ONLINE_OFFLINE_WAIT_TIME}
}


health_check() {
    exptime=0
    echo "checking ${HEALTH_CHECK_URL}"
    while true
    do
        status_code=`/usr/bin/curl -L -o /dev/null --connect-timeout 5 -s -w %{http_code}  ${HEALTH_CHECK_URL}`
        if [ x$status_code != x200 ];then
            sleep 1
            ((exptime++))
            echo -n -e "\rWait app to pass health check: $exptime..."
        else
            break
        fi
        if [ $exptime -gt ${APP_START_TIMEOUT} ]; then
            echo
            echo 'app start failed'
            exit 1
        fi
    done
    echo "check ${HEALTH_CHECK_URL} success"
}

start_application() {
    echo "start jar"
    if kill -0 "$(ps -ef | grep -v grep | grep -w [w]eimall-api-single | awk '{print $2}')"; then
        echo "Application is running, exit"
        exit 0
    fi
    cd ${APP_HOME}
    unzip -o -j ${JAR_NAME}  BOOT-INF/classes/application-prod.yml -d config/
    java -jar ${JAR_NAME} --spring.profiles.active=prod   -server  -XX:MetaspaceSize=512m -XX:MaxMetaspaceSize=512m -Xms2048m -Xmx4096m -Xmn256m -Xss256k -XX:SurvivorRatio=7 -XX:+UseConcMarkSweepGC -XX:+UseFastAccessorMethods  > /dev/null 2>&1   & 
}

stop_application() {
    echo "stop jar"
    pid=`ps -ef | grep -v grep | grep -w [w]eimall-api-single | awk '{print $2}'`
    kill ${pid}  ## 发送停止信号.
    sleep ${ONLINE_OFFLINE_WAIT_TIME}
    kill -9 ${pid}
    if [ $? -ne 0 ] ; then
	echo "stop failed "
    fi
}


start() {
    start_application
    health_check
    online
}

stop() {
    offline
    stop_application
}

case "$ACTION" in
    start)
        start
    ;;
    stop)
        stop
    ;;
    online)
        online
    ;;
    offline)
        offline
    ;;
    restart)
        stop
        start
    ;;
    *)
        usage
    ;;
esac
```

