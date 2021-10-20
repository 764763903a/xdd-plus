自用、学习、爱用不用。我就是上传数据也不叫偷，我自己脚本还不是想咋写就咋写，整天偷偷偷的，毛病。

## 注意事项

  1. master:的值即为密码，后面不可带注释，全匹配方可登录，也不要pt_pin 可自定义
   2. 2.9+版本需要配置    cid和secret 在青龙里面系统设置，添加应用后配置

# 支持作者

就当是支持支持作者更新吧 谢谢各位了。  几毛钱的包愿意就当支持我了 每天都可以领一次 建议保存到本地 辛苦了  不愿意也没事 几块钱的大家可以去买瓶水喝



<img src="https://i.loli.net/2021/10/11/odcbTfw4LM9svOS.jpg" alt="1.jpg" style="zoom:33%;" /><img src="https://i.loli.net/2021/10/11/ru73IykXez9g48K.jpg" alt="2.jpg" style="zoom: 23%;" />

# 更新日志

## 10-11

- 适配V2.8+版本包括即将来的3.0版本都适配了  = - 基本你们能用到黄了把
- 修复了QQ转账功能

## 10-07

- 新增延时设置，怕黑号调高，不怕的调低

## 10-04

- 新增QQ转账功能  格式 QQ转账 QQ 钱 例如        QQ转账  7647 100 
- 各位大老爷支持支持作者吧 谢谢啦

## 10-03

- 修复空pin可导入问题
- 剔除ck批量导入功能

## 10-02

- 修复删除指令没有删除jdcookiepool

## 10-01

- 大家安心过节 国庆快乐
- 修复WSkey自动无限更新问题
- 修复删除指令不够彻底问题

## 9-27

-  增加配置是否自动添加好友  注意看配置
- 增加短信自动绑定QQ （此接口面向傻妞对接）
- 增加lim配置，可对查询等各种命令进行次数限制。

- 会员中心代码内测中


## 9-24

- 即将上线CK提交页面，会员中心页面


## 9-21

- 新增加好友自动同意
- 修复清理过期账号指令
- 新增扫码登录（暂停）
- 新增删除WCK指令 删除过期wskey
- 修复之前临时屏蔽失效通知，临时关闭失效账号处理。
- 下个开发重点是登录页面和会员中心

## 9-20

感谢大佬增加xdd的接入

[Y佬的JD-qinglong](https://github.com/rubyangxg/jd-qinglong )

配置config中的apiToken，即可完成与Y佬的短信对接

XDD_URL格式为http://IP地址:端口/api/login/smslogin

![image-20210920125520315](C:\Users\76476\AppData\Roaming\Typora\typora-user-images\image-20210920125520315.png)

再次感谢Y佬  具体教程问问群里或者等好心群友提供我顶置更新

## 9-19

- 短信对接接口等待Y对接
- fix 七连失效通知

## 09-18之前

  1. 发送wskey即可自动添加账号
  2. 账号过期自动换key
  3. 定时十二小时自动换key
  4. 缓存token
  5. 批量绑定wskey
  6. 多容器 token缓存过期问题修复
  7. 解决%!(EXTRA 错误
  8. 手动指令更新
  9. 可替换失效wskey
  10. wskey失效检测
  11. Whiskey更新
  12. 新增删除账号指令
  13. 清理过期账号指令
  14. 合并详细查询功能
  15. 支持所有格式得CK  ALOOK  京东APP等啥都行
  16. 写入失败问题已解决wskey失效两次转换
  17. wskey过期提示
  18. 修复更新指定跳过空wskey
  19. 修复转换错误自动改为false 修复七次无限转换问题
  20. 新增 AtTime参数 不配置导致失败得别怪我
  21. AtTime:  #填写1-12之间的数  填错自负默认为10  10点容易出现高峰超时。
  22. IsHelp:   #填写true或者false  默认false 是否往容器添加助力码
  23. IsOldV4: #填写true或者false  默认false  是否新版或者旧版V4

## 重大更新

fix 重大BUG修复，解决以下几个问题，

- 新增账号部分参数空白
- 不管是CK新增还是WSKEY新增账号导致清空CK，由1导致的
- 新增添加后自动回复查询参数。
- 修复账号无限判错问题。
- 新增纯CK版本 可配置调整为WSKEY+CK  和纯CK版本
- Wskey: # 填空默认禁用wskey转换 需要的填true新增配置 默认关闭wskey 需要的自己设置下



## 开发目标

- plus登陆页面 
- 新增头部账号按次序进入容器助力模式
- 优化检测Ck代码
- 增加扫码登录，问题是只能QQ扫码登录，有点纠结要不要加




# 鸣谢

[jd-qinglong](https://github.com/rubyangxg/jd-qinglong )

[傻妞](https://github.com/rubyangxg/jd-qinglong )

[fuckee的拓展插件](https://github.com/ufuckee/jd_cookie)

以及大大小小的模块作者

# 安装教程 

xdd-plus安装教程
第一步：下载go
cd /usr/local && wget https://golang.google.cn/dl/go1.16.7.linux-amd64.tar.gz -O go1.16.7.linux-amd64.tar.gz
第二步：解压go
tar -xvzf go1.16.7.linux-amd64.tar.gz
第三步：设置环境变量 
vi /etc/profile
将文本复制到最后一行
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/path
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
第五步：先按   esc 
然后输入  :wq 
保存文件后 
source /etc/profile
第六步：检查go安装
go env
第七步：拉xdd-plus的库
cd ~ && git clone https://ghproxy.com/https://github.com/764763903a/xdd-plus.git
第八步：编译xdd-plus
cd /root/xdd-plus && go build

# 常见问题

编码问题参考
https://blog.csdn.net/qq_29499107/article/details/83583983
/usr/lib64/python3.6/http

Token故障请先用官方教程重装  已排查是nginx问题
https://thin-hill-428.notion.site/2-8Faker-QL-pannel-Faker-Repository-environment-Setup-45edcbfe90d74d8abb2d71896eab3be7
请使用官方一键安装 就解决此问题了



1. 如何自动更新短信镜像

```
docker run -d \
    --name watchtower \
    --restart always \
    -v /var/run/docker.sock:/var/run/docker.sock \
    containrrr/watchtower \
    --cleanup
```




有问题自己提需求啊。。。有空就解决没空靠自己了各位铁子

