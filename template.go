// template.go
package ssui

var HtmlPage1 = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>{{.Title}}</title>
  <link rel="icon" href="/uilib/images/favicon.ico">
  <link rel="stylesheet" href="/uilib/layui/css/layui.css" media="all">
  <script src="/uilib/layui/layui.js" charset="utf-8"></script>
  <script src="uilib/js/lay-config.js" charset="utf-8"></script>
</head>
<body>
<div class="layui-layout">
<div class="layui-container">
`

var HtmlPage2 = `
</div>
</div>
</body>
</html>
`

var HtmlHeader = `
	<form class="layui-form" action="">
	<div class="layui-row">
`

var HtmlFooter = `    	</div>
		</form>
</div>
</div>
`

var HtmlScript = `
<script>
layui.use(['form'], function () {
        var $ = layui.jquery,
            form = layui.form;

		form.on('checkbox(layui_checkbox)', function(data){
		  if(data.elem.checked){
				data.elem.value = '1';
		  }else{
				data.elem.value = '0';
		  }
		});
		form.render('select');
		form.render('radio');
		form.render('checkbox');
});

function isEmpty(obj){
    if(typeof obj == "undefined" || obj == null || obj == ""){
        return true;
    }else{
        return false;
    }
}

function buttonClick(e){
	var url = "/button_click?event_id="+e+"&"+getAllElemVal();
	var $ = layui.jquery;
	var loading = layer.load(0, {shade: false, time: 10 * 1000});
	$.get(url,function(ret){
		layer.close(loading);
		handleRsp(ret);
    });
}

function handleRsp(ret){
	var $ = layui.jquery;
	var layer = layui.layer;
	var obj = JSON.parse(ret);
	var miniPage = layui.miniPage;
 	console.log(obj);

	if(!isEmpty(obj.Error)){
		layer.msg(obj.Error);
	}
	

	if(obj.SelfClose){
		var index = parent.layer.getFrameIndex(window.name);
		parent.layer.close(index);
	}

	if(!isEmpty(obj.RedirectUrl))
	{
		if(!isEmpty(obj.ShowInDialog)){
			var index=layer.open({
			  	type: 2,
				title: obj.ShowInDialog,
				area: ['50%', '70%'], //宽高
			  	content: obj.RedirectUrl
			});
		}else{
			redirectUrl(obj.RedirectUrl);
		}
	}
}

function getAllElemVal(){
	var $ = layui.jquery;
`

var HtmlScriptFrame = `

function redirectUrl(url){
	var miniPage = layui.miniPage;
	var hashStr = window.location.hash.replace("#","");
	if(isEmpty(url)){
		miniPage.local();
		return;
	}else if(hashStr == url){
		miniPage.reload(url);
		return;
	}
	window.location.hash = url;
}
</script>
`

var HtmlScriptPage = `

function redirectUrl(url){
	if(isEmpty(url) || window.location.href==url){
		location.reload();
	}else{
		window.location.href=url;
	}
}
</script>`
