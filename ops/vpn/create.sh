#!/usr/bin/env bash

client=$1
group=$2

if [[ x$client = x || x$group = x ]]; then
    echo "Usage: $0 clientname group"
    exit 1
fi
cd /etc/openvpn/easy-rsa/3/

declare -A groupNetPrefixArr

// 对不同的部门进行分组
// 不同部门的权限不一样. 服务器统一管理
groupNetPrefixArr=(
  ["platform"]="172.30.0."
  ["spider"]="172.30.1."
  ["product"]="172.30.2."
  ["bi"]="172.30.10."
  ["data"]="172.30.100."
)
baseDir="/etc/openvpn"
clientConfigDir="$baseDir/junhsue/users"

_netPrefix=${groupNetPrefixArr[$group]}
netPrefix=${_netPrefix:? error group name!!}

while true
do
  randomNum=$(expr $RANDOM % 255)
  clientIp="${netPrefix}${randomNum}"
  if ! grep -qrn $clientIp $clientConfigDir ;then
      echo "-- user $client assign ip [$clientIp]"
      break
  fi
  echo "-- $clientIp already in use..."
  sleep 1s
done

echo "Generating profiles..."
cat > $clientConfigDir/$client << eof
ifconfig-push $clientIp 255.255.0.0
## >>$group<<
`cat $baseDir/junhsue/$group`
eof
echo "Generating done..."
#exit

if [[ ! -e keys/${client}.key ]]; then
    echo "Generating keys..."
    . vars
    ./easyrsa build-client-full $client nopass
    echo "...keys generated."
fi

tarball=./keys/$client.tgz

if [ ! -e $tarball ]; then
    echo "Creating tarball..."
    tmpdir=/tmp/client-tar.$$
    mkdir $tmpdir
    cp client.ovpn $tmpdir/client.ovpn
    sed -i "s/sample/$client/g" $tmpdir/client.ovpn
    cp keys/ca.crt $tmpdir
    cp pki/private/$client.key $tmpdir/
    cp pki/issued/$client.crt $tmpdir/
    tar -C $tmpdir -czvf $tarball .
    rm -rf $tmpdir
    echo "...tarball created"
else
    echo "Nothing to do, so nothing done. (keys/$client.tgz already exists)"
fi
