package models

import "fmt"

var Admin = `<html lang="zh-cn">

    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>账号管理</title>
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/cdle/static/layui/css/layui.css">
        <script src="https://cdn.jsdelivr.net/gh/cdle/static/layui/layui.all.js"></script>
    </head>
    
    <body>
        <div class="layui-tab">
            <ul class="layui-tab-title">
                <li class="layui-this">账号管理</li>
                <li>系统设置</li>
            </ul>
            <div class="layui-tab-content">
                <div class="layui-tab-item layui-show">
                    <table id="accounts" lay-filter="accounts"></table>
                </div>
                <div class="layui-tab-item">
                   啥都没有
                </div>
            </div>
    </body>
    <script>
        var table = layui.table;
        table.render({
            elem: '#accounts',
            height: "auto",
            url: '/api/account',
            toolbar: 'default',
            response: {
                statusName: 'code',
                statusCode: 200,
                msgName: 'code',
                countName: 'message',
                dataName: 'data'
            },
            title: '账号列表',
            page: true,
            limit: 15,
            cols: [
                [ //表头
                    {
                        field: 'ID',
                        title: 'ID',
                        width: 100,
                        align: 'center',
                    }, {
                        field: 'Nickname',
                        title: '用户昵称',
                        width: 120,
                        align: 'center',
                    }, {
                        field: 'CreateAt',
                        title: '创建时间',
                        edit: 'text',
                        width: 110,
                        align: 'center',
                    }, {
                        field: 'BeanNum',
                        title: '京豆数目',
                        width: 90,
                        align: 'center',
                    }, {
                        field: 'UserLevel',
                        title: '用户等级',
                        width: 90,
                        align: 'center',
                    }, {
                        field: 'LevelName',
                        title: '等级名称',
                        width: 100,
                        align: 'center',
                    }, {
                        field: 'Priority',
                        title: '优先级',
                        width: 80,
                        edit: 'text',
                        align: 'center',
                    }, {
                        field: 'Available',
                        title: '可用',
                        edit: 'text',
                        width: 80,
                        align: 'center',
                    }, {
                        field: 'Hack',
                        title: '屏蔽',
                        edit: 'text',
                        width: 80,
                        align: 'center',
                    }, {
                        field: 'Help',
                        title: '助力',
                        edit: 'text',
                        width: 80,
                        align: 'center',
                    }, {
                        field: 'Note',
                        title: '备注',
                        width: 120,
                        edit: 'text',
                        align: 'center',
                    }, {
                        field: 'PtPin',
                        title: 'PtPin',
                        width: 150,
                        align: 'center',
                    }, {
                        field: 'QQ',
                        title: 'QQ',
                        width: 120,
                        edit: 'text',
                        align: 'center',
                    }, {
                        field: 'PushPlus',
                        title: 'Push+',
                        width: 120,
                        edit: 'text',
                        align: 'center',
                    }, {
                        field: 'Fruit',
                        title: '东东农场',
                        edit: 'text',
                        width: 120,
                        align: 'center',
                    }, {
                        field: 'Pet',
                        title: '东东萌宠',
                        edit: 'text',
                        width: 120,
                        align: 'center',
                    }, {
                        field: 'Bean',
                        title: '种豆得豆',
                        edit: 'text',
                        width: 120,
                        align: 'center',
                    }, {
                        field: 'JdFactory',
                        title: '东东工厂',
                        edit: 'text',
                        width: 120,
                        align: 'center',
                    }, {
                        field: 'DreamFactory',
                        title: '惊喜工厂',
                        edit: 'text',
                        width: 120,
                        align: 'center',
                    }, {
                        field: 'Cash',
                        title: '签到领现金',
                        edit: 'text',
                        width: 120,
                        align: 'center',
                    }
                ]
            ]
        });
    
        table.on('edit(accounts)', function(obj) {
            obj.data.Priority = +obj.data.Priority
            obj.data.JinLi = +obj.data.JinLi
            obj.data.QQ = +obj.data.QQ
            layui.$.ajax({
                url: '/api/account',
                type: 'POST',
                contentType: "application/json",
                data: JSON.stringify(obj.data),
                dataType: 'json',
                timeout: 1000,
                cache: false,
                error: function() {
                    table.reload('accounts');
                }, //错误执行方法
                success: function(data) {
                    layer.msg(data["msg"])
                    table.reload('accounts');
                },
            });
        });
        table.on('toolbar(accounts)', function(obj){
            var checkStatus = table.checkStatus(obj.config.id);
            switch(obj.event){
              case 'add':
                layer.msg('添加');
              break;
              case 'delete':
                layer.msg('删除');
              break;
              case 'update':
                layer.msg('编辑');
              break;
            };
          });
    </script>
    
    
    </html>`

func Count() string {
	zs := 0
	yx := 0
	wx := 0
	tl := 0
	ts := 0
	tc := 0
	dt := Date()
	cks := GetJdCookies()
	for _, ck := range cks {
		zs++
		if ck.Available == True {
			yx++
		} else {
			wx++
		}
		if ck.CreateAt == dt {
			tc++
		}
	}
	jps := []JdCookiePool{}
	db.Find(&jps)
	for _, jp := range jps {
		if jp.CreateAt == dt {
			ts++
		}
		if jp.LoseAt == dt {
			tl++
		}
	}
	return fmt.Sprintf("总数%d,有效%d,无效%d,今日失效%d,今日扫码%d,今日新增%d", zs, yx, wx, tl, ts, tc)
}
