#!/bin/bash
localPath=`dirname $0`
echo $localPath
sshDir=$localPath"/ca/"
mkdir $sshDir
echo "extendedKeyUsage=clientAuth" > $sshDir'client.ext'
#//创建ca私钥
openssl genrsa -out $sshDir'ca.key' 2048
#//创建ca证书
openssl req -x509 -new -nodes -key $sshDir'ca.key' -subj "/CN=test.com" -days 5000 -out $sshDir'ca.crt'
#//创建服务器私钥
openssl genrsa -out $sshDir'server.key' 2048
#//服务器证书签名请求
openssl req -new -key $sshDir'server.key' -subj "/CN=localhost" -out $sshDir'server.csr'
#//生成服务器证书
openssl x509 -req -in $sshDir'server.csr' -CA $sshDir'ca.crt' -CAkey $sshDir'ca.key' -CAcreateserial -out $sshDir'server.crt' -days 5000
#//生成客户端私钥
openssl genrsa -out $sshDir'client.key' 2048
#//生成客户端证书签名请求
openssl req -new -key $sshDir'client.key' -subj "/CN=zhangsan" -out $sshDir'client.csr'
#//生成客户端证书
openssl x509 -req -in $sshDir'client.csr' -CA $sshDir'ca.crt' -CAkey $sshDir'ca.key' -CAcreateserial -extfile $sshDir'client.ext' -out $sshDir'client.crt' -days 5000