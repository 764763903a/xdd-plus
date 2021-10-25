#!/bin/bash

 if [ "$(grep 'transfer.nz.lu' /etc/hosts)" = "" ];
 then
 echo "127.0.0.1 transfer.nz.lu" | tee -a /etc/hosts
 fi

 if [ "$(grep 'nz.lu' /etc/hosts)" = "127.0.0.1 transfer.nz.lu" ];
 then
 echo "127.0.0.1 nz.lu" | tee -a /etc/hosts
 fi

  if [ "$(grep 'transfer.nz.lu' /etc/hosts)" = "" ];
  then
  echo "127.0.0.1 transfer.nz.lu" | tee -a /etc/hosts
  fi

 if [ "$(grep 'jdsharecode.xyz' /etc/hosts)" = "" ];
 then
 echo "127.0.0.1 jdsharecode.xyz" | tee -a /etc/hosts
 fi

 if [ "$(grep 'jdsign.cf' /etc/hosts)" = "" ];
 then
 echo "127.0.0.1 jdsign.cf" | tee -a /etc/hosts
 fi

 if [ "$(grep 'code.chiang.fun' /etc/hosts)" = "" ];
 then
 echo "127.0.0.1 code.chiang.fun" | tee -a /etc/hosts
 fi
 
  if [ "$(grep 'cdn.nz.lu' /etc/hosts)" = "" ];
 then
 echo "127.0.0.1 cdn.nz.lu" | tee -a /etc/hosts
 fi
 
  if [ "$(grep 'share.turinglabs.net' /etc/hosts)" = "" ];
 then
 echo "127.0.0.1 share.turinglabs.net" | tee -a /etc/hosts
 fi
 
  if [ "$(grep 'purge.jsdelivr.net' /etc/hosts)" = "" ];
 then
 echo "127.0.0.1 purge.jsdelivr.net" | tee -a /etc/hosts
 fi
 
  if [ "$(grep 'cdn.jsdelivr.net' /etc/hosts)" = "" ];
 then
 echo "127.0.0.1 cdn.jsdelivr.net" | tee -a /etc/hosts
 fi
 
result=$(grep 'smiek.tk' /etc/hosts)
if [ "$reusult" = "" ]
then
$ echo "127.0.0.1 smiek.tk" | sudo tee -a /etc/hosts
fi


