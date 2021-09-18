自用、学习、爱用不用。我就是上传数据也不叫偷，我自己脚本还不是想咋写就咋写，整天偷偷偷的，毛病。

## 注意事项

 1. master:的值即为密码，后面不可带注释，全匹配方可登录，也不要pt_pin 可自定义
  2. 2.9+版本需要配置    cid和secret 在青龙里面系统设置，添加应用后配置

# 更新日志

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
1.新增账号部分参数空白
2.不管是CK新增还是WSKEY新增账号导致清空CK，由1导致的
3.新增添加后自动回复查询参数。
4.修复账号无限判错问题。
新增纯CK版本 可配置调整为WSKEY+CK  和纯CK版本
Wskey: # 填空默认禁用wskey转换 需要的填true新增配置 默认关闭wskey
需要的自己设置下



## 开发目标

1.验证码登录直接接入

2.plus登陆页面 ---会员中心给你们看个样板图把



# 常见问题

编码问题参考
https://blog.csdn.net/qq_29499107/article/details/83583983
/usr/lib64/python3.6/http

Token故障请先用官方教程重装  已排查是nginx问题
https://thin-hill-428.notion.site/2-8Faker-QL-pannel-Faker-Repository-environment-Setup-45edcbfe90d74d8abb2d71896eab3be7
请使用官方一键安装 就解决此问题了




有问题自己提需求啊。。。有空就解决没空靠自己了各位铁子

