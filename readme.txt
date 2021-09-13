自用、学习、爱用不用。我就是上传数据也不叫偷，我自己脚本还不是想咋写就咋写，整天偷偷偷的，毛病。
##注意事项
 1. master: 的值即为密码，后面不可带注释，全匹配方可登录，也不要pt_pin 可自定义
 2. 2.9+版本需要配置    cid: admin
                  secret: admin
                 在青龙里面系统设置，添加应用后配置
 3.发送wskey即可自动添加账号
 4.账号过期自动换key
 5.定时十二小时自动换key
 6.缓存token
 7.批量绑定wskey
 8.多容器 token缓存过期问题修复
 10.解决%!(EXTRA 错误
11.手动指令更新
12.可替换失效wskey
13wskey失效检测
Whiskey更新
新增删除账号指令
清理过期账号指令
合并详细查询功能

14 找猫咪偷一下CK模糊检测
15 写入失败问题 不知道啥原因
16 wskey失效两次转换
17 wskey过期提示
全面适配所有CK格式  ALOOK  京东APP等啥都行
修复更新指定跳过空wskey
修复转换错误自动改为false 修复七次无限转换问题
新增 AtTime参数 不配置导致失败得别怪我
AtTime:  #填写1-12之间的数  填错自负默认为10  10点容易出现高峰超时。

以上是做完了 以下是待开发

新增纯CK版本 可配置调整为WSKEY+CK  和纯CK版本

考虑仓库私有化  另附如果不喜欢别用  ninja和xdd魔改版大有人做  免费开源没收你钱总不能还挨你一顿喷？

编码问题参考
https://blog.csdn.net/qq_29499107/article/details/83583983
/usr/lib64/python3.6/http

Token故障请先用官方教程重装  已排查是nginx问题
https://thin-hill-428.notion.site/2-8Faker-QL-pannel-Faker-Repository-environment-Setup-45edcbfe90d74d8abb2d71896eab3be7
请使用官方一键安装 就解决此问题了


全面适配新版V4  需要旧版V4的同学自己修改下 container代码 553行
	req := httplib.Post(c.Address + "/api/auth")
	修改为
	req := httplib.Post(c.Address + "/auth")




有问题自己提需求啊。。。有空就解决没空靠自己了各位铁子

