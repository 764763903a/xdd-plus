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
                    },{
                        field: 'WsKey',
                        title: 'WsKey',
                        width: 80,
                        edit: 'text',
                        align: 'center',
                    },{
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

var UserCenter = `
<!DOCTYPE html>
<!-- saved from url=(0043)http://v.bootstrapmb.com/2021/6/b2dr810290/ -->
<html style="font-size: 53.3333px;" class="hairlines">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>会员中心首页样式h5模板</title>
    <meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no" />
    <link rel="stylesheet" type="text/css" href="http://h5img.smxy.xyz/style1.css" />
    <script src="http://h5img.smxy.xyz/flexible.js" type="text/javascript" charset="utf-8"></script>
    <script src="http://h5img.smxy.xyz/zepto.min.js" type="text/javascript" charset="utf-8"></script>
    <style>.one-pan-tip { cursor: pointer;}.one-pan-tip::before {background-position: center;background-size: 100% 100%;background-repeat: no-repeat;box-sizing: border-box;width: 1em;height: 1em;margin: 0 1px .15em 1px;vertical-align: middle;display: inline-block;}.one-pan-tip-success::before {content: '';background-image: url(data:image/png;base64,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)}.one-pan-tip-error {text-decoration: line-through;}.one-pan-tip-error::before {content: '';background-image: url(data:image/png;base64,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)}.one-pan-tip-other::before {content: '';background-image: url(data:image/png;base64,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)}.one-pan-tip-lock::before{content: '';background-image: url(data:image/png;base64,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)}</style>
    <meta name="ljjc::status" content="on" />
    <style>.juejin-search[data-v-f493e070]{display:flex;width:682px;height:46px;border-radius:2px;flex-direction:row;align-items:center;justify-content:center;position:relative}.juejin-search .search-anim[data-v-f493e070]{position:absolute;left:8px;width:28px;height:28px;object-fit:contain;animation-play-state:paused}.juejin-search .search-anim.slide-right-enter-active[data-v-f493e070],.juejin-search .search-anim.slide-right-leave-active[data-v-f493e070]{transition:width .3s linear}.juejin-search .search-anim.slide-right-enter-from[data-v-f493e070],.juejin-search .search-anim.slide-right-leave-to[data-v-f493e070]{width:0}.juejin-search .juejin-search-logo[data-v-f493e070]{right:16px;position:absolute;width:23px;height:18px;object-fit:contain}.juejin-search .juejin-search-logo path[data-v-f493e070]{transition:all .3s linear}.juejin-search #juejin-search-input-global.input[data-v-f493e070]{padding:0 39px 0 33px;width:100%;height:100%;outline:0;border:none;border-radius:2px;background-color:#fff;color:#1d2129;font-size:18px;line-height:22px;font-weight:500;caret-color:transparent;box-sizing:border-box;background:rgba(148,191,255,.1)}.juejin-search #juejin-search-input-global.input.active[data-v-f493e070]{border:2px solid #bedaff}.juejin-search #juejin-search-input-global.input.animation-stopped[data-v-f493e070]{caret-color:#1e80ff;padding-left:16px}.juejin-search #juejin-search-input-global.input[data-v-f493e070]::placeholder{font-weight:400;color:#86909c}:root{--color-input-bg:#f4f5f5;--color-input-error-bg:#ffece8;--color-input-placeholder:#86909c;--color-input-text:#4e5969;--color-input-icon:#f53f3f}:root .dark{--color-input-bg:rgba(255, 255, 255, 0.12);--color-input-error-bg:rgba(255, 81, 50, 0.15);--color-input-placeholder:#e3e3e3;--color-input-text:#e3e3e3;--color-input-icon:#ff6247}[data-v-4c531118]:root{--color-brand:#1E80FF;--color-brand-light:#E8F3FF;--color-nav-title:#86909C;--color-nav-popup-bg:#FFFFFF;--color-primary:#1D2129;--color-secondary:#4E5969;--color-thirdly:#86909C;--color-hover:#1e80ff;--color-hover-thirdly:#86909c;--color-dropdown-text:#1E80FF;--color-divider:#E5E6EB;--color-main-bg:#f4f5f5;--color-secondary-bg:#FFFFFF;--color-thirdly-bg:#f4f5f5;--color-hover-bg:#E8F3FF;--color-comment-bg:rgba(244, 245, 245, 0.5);--hover-bg:linear-gradient(
            90deg,
            rgba(232, 243, 255, 0) 0%,
            rgba(232, 243, 255, 0.8) 25.09%,
            #e8f3ff 50.16%,
            rgba(232, 243, 255, 0.8) 75.47%,
            rgba(232, 243, 255, 0) 100%
    );--color-mask:rgba(0, 0, 0, 0.4);--color-quick-nav-text:#ffffff;--color-nav-bg:rgba(255, 255, 255, 0.13);--color-nav-selected-border:rgba(229, 230, 235, 0.3);--color-tips:#F53F3F;--color-fourthly:#C9CDD4;--color-shadow:rgba(0, 0, 0, 0.16);--color-grey-triangle:#e5e6eb;--color-icon-search:#ffffff;--color-navbar-icon:#1e80ff;--color-layout-dropdown-bg:rgba(232, 243, 255, 0.8);--color-layout-title:#4E5969;--color-layout-title-active:#1E80FF;--color-layout-icon-outline:rgba(30, 128, 255, 0.5);--color-layout-icon-fill:#ffffff}:root .dark[data-v-4c531118]{--color-brand:#1352a3;--color-nav-title:#e3e3e3;--color-nav-popup-bg:#1352A3;--color-primary:#e3e3e3;--color-secondary:#a9a9a9;--color-thirdly:#7d7d7f;--color-hover:#eeeeee;--color-hover-thirdly:#878789;--color-dropdown-text:#878789;--color-divider:#4a4a4a;--color-main-bg:#121212;--color-secondary-bg:#272727;--color-thirdly-bg:#3a3a3a;--color-hover-bg:#3a3a3a;--color-comment-bg:#313131;--hover-bg:linear-gradient(
            90deg,
            rgba(58, 58, 58, 0) 0%,
            rgba(58, 58, 58, 0.8) 25.09%,
            #3a3a3a 50.16%,
            rgba(58, 58, 58, 0.8) 75.47%,
            rgba(58, 58, 58, 0) 100%
    );--color-mask:rgba(0, 0, 0, 0.4);--color-quick-nav-text:#e3e3e3;--color-nav-bg:rgb(30, 30, 30);--color-nav-selected-border:#4a4a4a;--color-tips:#bc3030;--color-fourthly:#878789;--color-shadow:rgba(0, 0, 0, 0.16);--color-grey-triangle:#3a3a3a;--color-icon-search:#e3e3e3;--color-navbar-icon:#e3e3e3;--color-layout-dropdown-bg:rgba(125, 125, 127, 0.8);--color-layout-title:#eeeeee;--color-layout-title-active:#eeeeee;--color-layout-icon-outline:#131313;--color-layout-icon-fill:#e3e3e3}.input-option[data-v-4c531118]{display:flex;flex-direction:column}.input-option span.error[data-v-4c531118]{margin-left:5.1666666667rem;font-size:1rem;line-height:20px;display:inline-block;height:20px;color:var(--color-tips)}.input-wrapper[data-v-4c531118]{display:flex;flex-direction:row;align-items:center;width:100%}.input-wrapper label[data-v-4c531118]{width:4em;font-size:1.1666666667rem;line-height:1.8333333333rem;color:var(--color-thirdly);margin-right:1rem}.input-wrapper .input[data-v-4c531118]{flex:1 0 auto;position:relative}.input-wrapper .input.error .input-inner[data-v-4c531118]{background-color:var(--color-input-error-bg)}.input-wrapper .input.error .btn-clear[data-v-4c531118]{color:var(--color-input-icon)}.input-wrapper .input .input-inner[data-v-4c531118]{background:var(--color-input-bg);border-radius:2px;color:var(--color-input-text);font-size:1.0833333333rem;line-height:1.8333333333rem;height:2.3333333333rem;padding:0 8px;outline:0;border:none;width:100%}.input-wrapper .input .input-inner[data-v-4c531118]::placeholder{color:var(--color-input-placeholder)}.input-wrapper .btn-clear[data-v-4c531118]{position:absolute;top:50%;right:0;transform:translateY(-50%);background:0 0;border:none;outline:0;color:var(--color-fourthly)}.input-wrapper .btn-clear[data-v-4c531118]::before{font-size:10px;line-height:10px}[data-v-25df73b2]{box-sizing:border-box}.color-tool[data-v-25df73b2]{padding:0 16px!important}.color-tool .row[data-v-25df73b2]{display:flex;align-items:center}.color-tool .color-picker[data-v-25df73b2]{cursor:pointer;outline:0;border:none;padding:0;margin:0;border-radius:2px;background-color:transparent;width:92px;height:40px}.color-tool .color-picker[data-v-25df73b2]::-webkit-color-swatch-wrapper{padding:3px;border:1px solid transparent;border-radius:4px;transition:all .15s linear}.color-tool .color-picker[data-v-25df73b2]::-webkit-color-swatch-wrapper:hover{border:1px solid #bedaff}.color-tool .color-picker[data-v-25df73b2]::-webkit-color-swatch{border-radius:2px;border:none}.color-tool .input[data-v-25df73b2]{transform:translateY(10px);flex:1 1 auto;margin:0 12px}.color-tool .input[data-v-25df73b2] input.input-inner{height:40px;padding-left:16px;font-size:14px;color:#1d2129;box-sizing:border-box;background:#f4f5f5}.color-tool .input[data-v-25df73b2] label{display:none}.color-tool .input[data-v-25df73b2] span.error{margin-left:16px}.color-tool .input[data-v-25df73b2] .input-wrapper .btn-clear{right:8px}.color-tool .input[data-v-25df73b2] .input-wrapper .btn-clear::before{font-size:14px;color:#c9cdd4}.color-tool button[data-v-25df73b2]{outline:0;border:none;background-color:unset;width:93px;height:40px;font-size:14px}.color-tool .btn-convert[data-v-25df73b2]{background:#1e80ff;border-radius:2px;color:#fff;transition:all .15s linear}.color-tool .btn-convert[data-v-25df73b2]:hover{background:#5399ff}.color-tool .btn-convert[data-v-25df73b2]:active{background:#0060dd}.color-tool .btn-copy[data-v-25df73b2]{background:rgba(30,128,255,.05);border:1px solid rgba(30,128,255,.3);border-radius:2px;color:#1e80ff;transition:all .15s linear}.color-tool .btn-copy[data-v-25df73b2]:hover{background:rgba(30,128,255,.1);border-color:rgba(30,128,255,.45)}.color-tool .btn-copy[data-v-25df73b2]:active{background:rgba(30,128,255,.2);border-color:rgba(30,128,255,.6)}.color-tool .display[data-v-25df73b2]{flex:1;text-align:start;background-color:#f4f5f5;height:40px;margin:0 12px;border-radius:2px;line-height:40px;padding-left:16px;font-size:14px;color:#1d2129}.color-tool .label[data-v-25df73b2]{width:92px;font-size:16px;font-weight:500;color:#1d2129;text-align:end}.color-tool .row[data-v-25df73b2]:not(:first-of-type){margin-top:16px}.tool[data-v-5c9b7424]{width:100%;height:100%}iframe[data-v-5c9b7424]{min-height:488px}.calculator[data-v-81e152a8]{display:flex;align-items:center;padding:14px 16px 14px 24px}.calculator .result[data-v-81e152a8]{font-size:16px;font-weight:500;line-height:22px;color:#1d2129;margin:0 16px;text-overflow:ellipsis;flex:1 0 auto;overflow:hidden;white-space:nowrap;max-width:490px}.calculator .hint[data-v-81e152a8]{font-size:14px;line-height:22px;color:#86909c}.search-action[data-v-71378e58]{display:flex;align-items:center;padding:0 16px 0 20px;box-sizing:border-box;user-select:none;cursor:pointer;height:44px;border-left:4px solid transparent;border-top:4px solid transparent;border-bottom:4px solid transparent;transition:all .15s linear}.search-action.active[data-v-71378e58],.search-action[data-v-71378e58]:hover{border-left-color:#1e80ff;background-color:#f4f5f5}.search-action .search-content[data-v-71378e58]{display:flex;align-items:center;flex:1 0 auto;margin-right:16px}.search-action .search-content .search-content__logo[data-v-71378e58]{width:32px;height:32px}.search-action .search-content .search-content__engine[data-v-71378e58],.search-action .search-content .search-content__keyword[data-v-71378e58]{font-size:16px;font-weight:500;line-height:22px}.search-action .search-content .search-content__keyword[data-v-71378e58]{color:#1d2129;margin:0 4px 0 16px;text-overflow:ellipsis;overflow:hidden;white-space:nowrap;max-width:396px}.search-action .search-content .search-content__engine[data-v-71378e58]{color:#1e80ff}.search-action .hint[data-v-71378e58]{font-size:14px;line-height:22px;color:#1e80ff}[data-v-b7831bf2]:root{--color-brand:#1E80FF;--color-brand-light:#E8F3FF;--color-nav-title:#86909C;--color-nav-popup-bg:#FFFFFF;--color-primary:#1D2129;--color-secondary:#4E5969;--color-thirdly:#86909C;--color-hover:#1e80ff;--color-hover-thirdly:#86909c;--color-dropdown-text:#1E80FF;--color-divider:#E5E6EB;--color-main-bg:#f4f5f5;--color-secondary-bg:#FFFFFF;--color-thirdly-bg:#f4f5f5;--color-hover-bg:#E8F3FF;--color-comment-bg:rgba(244, 245, 245, 0.5);--hover-bg:linear-gradient(
            90deg,
            rgba(232, 243, 255, 0) 0%,
            rgba(232, 243, 255, 0.8) 25.09%,
            #e8f3ff 50.16%,
            rgba(232, 243, 255, 0.8) 75.47%,
            rgba(232, 243, 255, 0) 100%
    );--color-mask:rgba(0, 0, 0, 0.4);--color-quick-nav-text:#ffffff;--color-nav-bg:rgba(255, 255, 255, 0.13);--color-nav-selected-border:rgba(229, 230, 235, 0.3);--color-tips:#F53F3F;--color-fourthly:#C9CDD4;--color-shadow:rgba(0, 0, 0, 0.16);--color-grey-triangle:#e5e6eb;--color-icon-search:#ffffff;--color-navbar-icon:#1e80ff;--color-layout-dropdown-bg:rgba(232, 243, 255, 0.8);--color-layout-title:#4E5969;--color-layout-title-active:#1E80FF;--color-layout-icon-outline:rgba(30, 128, 255, 0.5);--color-layout-icon-fill:#ffffff}:root .dark[data-v-b7831bf2]{--color-brand:#1352a3;--color-nav-title:#e3e3e3;--color-nav-popup-bg:#1352A3;--color-primary:#e3e3e3;--color-secondary:#a9a9a9;--color-thirdly:#7d7d7f;--color-hover:#eeeeee;--color-hover-thirdly:#878789;--color-dropdown-text:#878789;--color-divider:#4a4a4a;--color-main-bg:#121212;--color-secondary-bg:#272727;--color-thirdly-bg:#3a3a3a;--color-hover-bg:#3a3a3a;--color-comment-bg:#313131;--hover-bg:linear-gradient(
            90deg,
            rgba(58, 58, 58, 0) 0%,
            rgba(58, 58, 58, 0.8) 25.09%,
            #3a3a3a 50.16%,
            rgba(58, 58, 58, 0.8) 75.47%,
            rgba(58, 58, 58, 0) 100%
    );--color-mask:rgba(0, 0, 0, 0.4);--color-quick-nav-text:#e3e3e3;--color-nav-bg:rgb(30, 30, 30);--color-nav-selected-border:#4a4a4a;--color-tips:#bc3030;--color-fourthly:#878789;--color-shadow:rgba(0, 0, 0, 0.16);--color-grey-triangle:#3a3a3a;--color-icon-search:#e3e3e3;--color-navbar-icon:#e3e3e3;--color-layout-dropdown-bg:rgba(125, 125, 127, 0.8);--color-layout-title:#eeeeee;--color-layout-title-active:#eeeeee;--color-layout-icon-outline:#131313;--color-layout-icon-fill:#e3e3e3}.search-app[data-v-b7831bf2]{z-index:9999;padding-top:160px;position:fixed;left:0;right:0;top:0;bottom:0;display:flex;align-items:flex-start;justify-content:center}.search-app.extension[data-v-b7831bf2]{z-index:700}@media (max-height:720px){.search-app.tool-active[data-v-b7831bf2]{padding-top:80px}}@media (max-height:640px){.search-app.tool-active[data-v-b7831bf2]{padding-top:30px}}.search-app .search-app__wrapper__[data-v-b7831bf2]{border-radius:2px;border:1px solid #1e80ff;background:#fff;box-shadow:0 0 0 4px rgba(30,128,255,.2),0 0 20px rgba(0,0,0,.15);backdrop-filter:blur(15px)}.search-app .search-app__wrapper__ .search-result[data-v-b7831bf2]{margin-top:8px}.search-app .search-app__wrapper__ .search-result .tool[data-v-b7831bf2]{padding:0 8px}.search-app .search-app__wrapper__ .search-result .setting-hint[data-v-b7831bf2]{display:flex;align-items:center;justify-content:flex-end;margin:0 16px;padding:12px 0 16px}.search-app .search-app__wrapper__ .search-result .setting-hint .text[data-v-b7831bf2]{color:#86909c;line-height:22px;cursor:pointer;user-select:none}.search-app .search-app__wrapper__ .search-result .setting-hint .text[data-v-b7831bf2]:hover:not(.disabled){color:#1e80ff;transition:all .15s linear}.search-app .search-app__wrapper__ .search-result .setting-hint .text.disabled[data-v-b7831bf2]{cursor:initial}.search-app .juejin-search[data-v-b7831bf2]{margin:8px}</style>
</head>
<body style="font-size: 24px;">
<!--返回-->
<div class="top">
    <i></i>会员中心
</div>
<!--个人信息-->
<div class="info-card">
    <div class="user">
        <img src="http://h5img.smxy.xyz/head.png" />
        <div class="name">
            <p>15555952_p</p>
            <span>升级至尊会员享额外特权</span>
        </div>
    </div>
    <div class="info">
        <div class="item">
            <p>优先级</p>
            <span>100</span>
        </div>
        <div class="item">
            <p>京豆</p>
            <span>4375</span>
        </div>
        <div class="item">
            <p>红包</p>
            <span>26.31</span>
        </div>
    </div>
</div>
<!--会员充值-->
<h5 class="title">会员升级</h5>
<div class="hycz">
    <ul>
        <li class="active">
            <div class="desc">
                <p class="type">12个月</p>
                <p class="money"><span>￥</span>66</p>
                <div class="remark">
                    <p>到期自动续费</p>
                    <p>可随时关闭</p>
                </div>
            </div>
            <div class="limit">
                限时特价 送豪礼
            </div></li>
        <li>
            <div class="desc">
                <p class="type">1个月</p>
                <p class="money"><span>￥</span>6</p>
                <div class="remark">
                    <p>联合月卡</p>
                </div>
            </div></li>
        <li>
            <div class="desc">
                <p class="type">连续包月</p>
                <p class="money"><span>￥</span>5</p>
                <div class="remark">
                    <p>到期自动续费</p>
                    <p>可随时关闭</p>
                </div>
            </div></li>
        <li>
            <div class="desc">
                <p class="type">连续包年</p>
                <p class="money"><span>￥</span>55</p>
                <div class="remark">
                    <p>到期自动续费</p>
                    <p>可随时关闭</p>
                </div>
            </div></li>
        <div class="holder"></div>
    </ul>
    <div class="btn">
        <button id="pay">立即升级</button>
        <p>升级立得800京豆</p>
    </div>
</div>
<h5 class="title">会员特权</h5>
<div class="hytq">
    <ul>
        <li><img src="http://h5img.smxy.xyz/icon-jinbi.png" /><p>每日京豆</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-shenjuan.png" /><p>京东农场</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-fanli.png" /><p>东东工厂</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-shangpin.png" /><p>京喜牧场</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-jisu.png" /><p>京东工厂</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-fenxiang.png" /><p>免费水果</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-huodong.png" /><p>专享活动</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-shengri.png" /><p>生日折扣</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-tixing.png" /><p>兑换提醒</p></li>
        <li><img src="http://h5img.smxy.xyz/icon-kefu.png" /><p>专属客服</p></li>
    </ul>
</div>
<div style="display:none">
    <a href="http://www.bootstrapmb.com/" one-link-mark="yes">更多前端代码</a>
</div>
<script type="text/javascript">$('.hycz ul li').click(function(){var index = $(this).index();$(this).addClass('active').siblings().removeClass('active');$('.hycz ul').scrollLeft($(this).width() * index)}) $('#pay').click(function(){var type = $('.hycz ul li.active').find('.type').text();var money = $('.hycz ul li.active').find('.money').text();confirm("确定要充值"+type+"会员吗？"+"价值："+money.substr(1)+"元")}) </script>
<!--下面是无用代码-->
</body>
</html>
`

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
