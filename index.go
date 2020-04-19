// index.go
package ssui

var IndexHtml = `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>%s</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta http-equiv="Access-Control-Allow-Origin" content="*">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="format-detection" content="telephone=no">
    <link rel="icon" href="/uilib/images/favicon.ico">
    <link rel="stylesheet" href="/uilib/layui/css/layui.css" media="all">
    <link rel="stylesheet" href="/uilib/css/layuimini.css?v=2.0.0" media="all">
    <link rel="stylesheet" href="/uilib/css/public.css" media="all">
    <style id="layuimini-bg-color">
    </style>
</head>
<body class="layui-layout-body layuimini-all">
<div class="layui-layout layui-layout-admin">

    <div class="layui-header header">
        <div class="layui-logo layuimini-logo layuimini-back-home"></div>

        <div class="layuimini-header-content">
            <a>
                <div class="layuimini-tool"><i title="展开" class="layui-icon layui-icon-shrink-right" data-side-fold="1"></i></div>
            </a>

            <!--电脑端头部菜单-->
            <ul class="layui-nav layui-layout-left layuimini-header-menu mobile layui-hide-xs layuimini-menu-header-pc">
            </ul>

            <!--手机端头部菜单-->
            <ul class="layui-nav layui-layout-left layuimini-header-menu mobile layui-hide-sm">
                <li class="layui-nav-item">
                    <a href="javascript:;"><i class="layui-icon layui-icon-template-1"></i> 选择模块</a>
                    <dl class="layui-nav-child layuimini-menu-header-mobile">
                    </dl>
                </li>
            </ul>

            <ul class="layui-nav layui-layout-right">

                <li class="layui-nav-item" lay-unselect>
                    <a href="javascript:;" data-refresh="刷新"><i class="layui-icon layui-icon-refresh"></i></a>
                </li>
                <li class="layui-nav-item" lay-unselect>
                    <a href="javascript:;" data-clear="清理" class="layuimini-clear"><i class="layui-icon layui-icon-delete"></i></a>
                </li>
                <li class="layui-nav-item mobile layui-hide-xs" lay-unselect>
                    <a href="javascript:;" data-check-screen="full"><i class="layui-icon layui-icon-screen-full"></i></a>
                </li>
                <li class="layui-nav-item layuimini-setting">
                    <a href="javascript:;">%s</a>
                    <dl class="layui-nav-child">
                        <dd>
                            <a href="javascript:;" layuimini-content-href="/usersetting" data-title="基本资料" data-icon="layui-icon layui-icon-username">基本资料</a>
                        </dd>
                        <dd>
                            <a href="javascript:;" layuimini-content-href="/chpwd" data-title="修改密码" data-icon="layui-icon layui-icon-password">修改密码</a>
                        </dd>
                        <dd>
                            <hr>
                        </dd>
                        <dd>
                            <a href="javascript:;" class="login-out">退出登录</a>
                        </dd>
                    </dl>
                </li>
                <li class="layui-nav-item layuimini-select-bgcolor mobile layui-hide-xs" lay-unselect>
                    <a href="javascript:;" data-bgcolor="配色方案"><i class="layui-icon layui-icon-more-vertical"></i></a>
                </li>
            </ul>
        </div>
    </div>

    <!--无限极左侧菜单-->
    <div class="layui-side layui-bg-black layuimini-menu-left">
    </div>

    <!--初始化加载层-->
    <div class="layuimini-loader">
        <div class="layuimini-loader-inner"></div>
    </div>

    <!--手机端遮罩层-->
    <div class="layuimini-make"></div>

    <!-- 移动导航 -->
    <div class="layuimini-site-mobile"><i class="layui-icon"></i></div>

    <div class="layui-body">
	<div class="layuimini-content-page">
	<div class="layuimini-container layuimini-page-anim">
    <div class="layuimini-main">
        <div class="layuimini-content">
        </div>
	</div>
    </div>
	</div>
    </div>

</div>
<script src="uilib/layui/layui.js" charset="utf-8"></script>
<script src="uilib/js/lay-config.js?v=2.0.0" charset="utf-8"></script>
<script>
    layui.use(['jquery', 'layer', 'miniAdmin'], function () {
        var $ = layui.jquery,
            layer = layui.layer,
            miniAdmin = layui.miniAdmin;

        var options = {
            iniUrl: "api/init",    // 初始化接口
            clearUrl: "api/clear", // 缓存清理接口
            renderPageVersion: true,    // 初始化页面是否加版本号
            bgColorDefault: 0,          // 主题默认配置
            multiModule: true,          // 是否开启多模块
            menuChildOpen: false,       // 是否默认展开菜单
            loadingTime: 0,             // 初始化加载时间
            pageAnim: true,             // 切换菜单动画
        };
        miniAdmin.render(options);

        $('.login-out').on("click", function () {
			//清除cookie
			var keys=document.cookie.match(/[^ =;]+(?=\=)/g);
            if (keys)
            {
                for (var i = keys.length; i--;)
                {
                	console.log(keys[i]);
                    document.cookie=keys[i]+'=0;expires=' + new Date(0).toUTCString()+";path=/";
                }
            }
            layer.msg('退出登录成功');
            window.location.href = 'login';
        });
    });
</script>
</body>
</html>
`
