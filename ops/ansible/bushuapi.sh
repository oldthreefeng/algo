#!/bin/bash

project="api"
current_version=$1

function dingding() {
  title="$project-$current_version"
  messageUrl="https://mall.youpenglai.com/apis/version"
  picUrl="http://icons.iconarchive.com/icons/paomedia/small-n-flat/1024/sign-check-icon.png"
  text="部署成功"
  PHONE="1582156468"
  TOKEN=$1
  DING="curl -H \"Content-Type: application/json\" -X POST --data '{\"msgtype\": \"link\", \"link\": {\"messageUrl\": \"${messageUrl}\", \"title\": \"${title}\", \"picUrl\": \"${picUrl}\", \"text\": \"${text}\",}, \"at\": {\"atMobiles\": [${PHONE}], \"isAtAll\": false}}' ${TOKEN}"
  eval "$DING"
}

if [[ x"$current_version" == x ]]; then
	echo -e "api 项目发布脚本\n\n\tsh $0  [version]"
	echo "example: sh $0 2.2.0.1"
	exit
fi

ansible-playbook api.yml  -e version="$current_version" -f 1

if [[ $? -eq 0 ]]; then
	dingding "https://oapi.dingtalk.com/robot/send?access_token=01bc245b59a337090fca147c123488de188d00cc56e60c77c3c573ddfae655b9"
fi