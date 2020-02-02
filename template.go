// template.go
package ssui

var HtmlHeader = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>{{.Title}}</title>
  <link rel="stylesheet" href="/layui/css/layui.css">
  <script src="/layui/layui.js"></script>
</head>
<body>

<div class="layui-footer" style="text-align:center;">
    <!-- 顶部固定区域 -->
    <h1>{{.Header}}</h1>
</div>

<div class="layui-container">
<div class="layui-main">
<form class="layui-form layui-form-pane" action="">
`

var HtmlFooter = `
<pre class="layui-code" id="retcode" ></pre>
</form>

</div>
</div>

<div class="layui-footer" style="text-align:center;">
    <!-- 底部固定区域 -->
    <span>{{.Footer}}</span>

</div>
</body>
</html>`

var HtmlScript1 = `<script>
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
});

function isEmpty(obj){
    if(typeof obj == "undefined" || obj == null || obj == ""){
        return true;
    }else{
        return false;
    }
}

function tableDel(tid,rid){
	var layer = layui.layer;
	layer.msg('Delete?', {
	  time: 0 //不自动关闭
	  ,btn: ['OK', 'Cancel']
	  ,yes: function(index){
	    layer.close(index);
	    var url = "/table_del?event_id="+tid+"&url_router="+window.location.pathname+"&rowid="+rid;
		var $ = layui.jquery;
		$.get(url,function(ret){
			handleRsp(ret);
	    });
	  }
	});
}

function buttonClick(e){
	var url = "/button_click?event_id="+e+"&url_router="+window.location.pathname;
	var param = getAllElemVal();
	if(!isEmpty(param))
	{
		url=url+"&"+param
	};
	var $ = layui.jquery;
	$.get(url,function(ret){
		handleRsp(ret);
    });
}

function handleRsp(ret){
	var $ = layui.jquery;
	var layer = layui.layer;
	var obj = JSON.parse(ret);
 	console.log(obj)
	if(!isEmpty(obj.Error)){
		 layer.msg(obj.Error);
	}else if(!isEmpty(obj.Content)){
		 $('#retcode').html(obj.Content+"<br>"+$('#retcode').html());
	}else if(obj.ShowInDialog){
		layer.open({
		  type: 2,
		  area: ['700px', '450px'],
		  fixed: false, //不固定
		  maxmin: true,
		  content: obj.Url
		});
	}else{
		redirectUrl(obj.Url)
	}
}

function redirectUrl(url){
	if(isEmpty(url)){
		location.reload();
	}else{
		window.location.href=url;
	}
}

function getAllElemVal(){
	var $ = layui.jquery;
`

var HtmlScript2 = `
</script>`

func SliceElem(s []string, i int) string {
	if len(s) == 0 {
		return ""
	}
	if i >= len(s) {
		i = 0
	}
	return s[i]
}
