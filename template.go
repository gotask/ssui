// template.go
package ssui

var HtmlHeader = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>{{.Title}}</title>
  <link rel="shortcut icon" type="image/ico" href="/layui/favicon.ico" />
  <link rel="stylesheet" href="/layui/css/layui.css">
  <script src="/layui/layui.js"></script>
</head>
<body>

<div class="layui-header" style="text-align:center;">
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
	    var url = "/table_del?event_id="+tid+"&token="+getToken()+"&url_router="+window.location.pathname+"&rowid="+rid;
		var $ = layui.jquery;
		$.get(url,function(ret){
			handleRsp(ret);
	    });
	  }
	});
}

function tableCpy(tid,ridx){
	var table = document.getElementById(tid);
	var child = table.getElementsByTagName("tr")[ridx + 1];
	var cols = child.getElementsByTagName("td");
	for(j = 0; j < cols.length-1; j++) {
	   document.getElementById(tid+"_"+j).value=cols[j].innerText;
	} 
}

function buttonClick(e){
	var url = "/button_click?event_id="+e+"&url_router="+window.location.pathname+"&"+getAllElemVal();
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
	if(!isEmpty(obj.Token)){
		 setToken(obj.Token);
	}
	if(!isEmpty(obj.Error)){
		 layer.msg(obj.Error);
	}else if(!isEmpty(obj.EchoText)){
		 $('#retcode').html(obj.EchoText+"<br>"+$('#retcode').html());
	}else if(obj.ShowInDialog){
		layer.open({
		  type: 2,
		  area: ['700px', '450px'],
		  fixed: false, //不固定
		  maxmin: true,
		  content: obj.RedirectUrl
		});
	}else{
		redirectUrl(obj.RedirectUrl)
	}
}

function getToken(){
	var usertoken = layui.data('userdata').token;
	if(isEmpty(usertoken)){
		return ""
	}
	return usertoken
}

function setToken(token){
	if(isEmpty(token)){
		return
	}
	layui.data('userdata', {
        key: 'token'
        ,value: token
    });
}
function delToken(){
	layui.data('userdata', {
        key: 'token'
        ,remove: true
    });
}

function redirectUrl(url){
	if(isEmpty(url)){
		location.reload();
	}else{
		var usertoken = layui.data('userdata').token;
		if(isEmpty(usertoken)){
			window.location.href=url
			return
		}
		if(url.indexOf("?")<0){
			url=url+"?"
		}else{
			url=url+"&"
		}
		window.location.href=url+"token="+getToken();
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
