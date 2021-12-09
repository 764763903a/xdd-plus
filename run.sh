#!/bin/bash
arch=`uname -m`
case $arch in
x86_64)
     arch="amd64"
     ;;
aarch64)
     arch="arm64"
     ;;
*)
     arch="arm"
     ;;
esac
filename="xdd-linux-${arch}"
url="https://github.91chi.fun/https://github.com/764763903a/xdd-plus/releases/download/main/${filename}"
dirname="xdd"
cd $HOME
if [ ! -d dirname ];then
  mkdir dirname
fi
cd xdd
#curl -L $url -O $filename
curl -L $url -o xdd
chmod 777 xdd
./xdd -d
