#!/usr/bin/env bash
source vars
cd /etc/openvpn/easy-rsa/3/
client=$1
./easyrsa revoke $client
./easyrsa gen-crl
find ./ -name "$client*" -exec rm -fv {} \;
