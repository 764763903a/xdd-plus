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
filename="xdd_linux_${arch}"
url="https://ghproxy.com/https://github.com/cdle/jd_study/releases/download/main/${filename}"
dirname="xdd"
cd $HOME
if [ ! -d dirname ];then
  mkdir dirname
fi
cd xdd
curl -L $url -O $filename
