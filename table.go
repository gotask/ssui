// table.go
package ssui

import (
	"strings"
)

type TableOperType int

var (
	TOEdit TableOperType = 1
	TOAdd  TableOperType = 2
	TODel  TableOperType = 3
)

type TableColumnType int

var (
	TCNormal TableColumnType = 0
	TCImg    TableColumnType = 1
	TCUrl    TableColumnType = 2
)

var GlobalTables = make(map[string]*HTable, 0)

func GetTable(id string) *HTable {
	if t, ok := GlobalTables[id]; ok {
		return t
	}
	return nil
}

type OnTableGetData func(user string, page, limit int, searchtxt string) (total int, data [][]string)
type OnTableEvent func(user string, t TableOperType, cols []string) ApiRsp
type OnTableUrl func(user, href string) string

type HTable struct {
	*ElemBase
	Header  []string
	ColType []int
	Tool    bool
	Search  bool

	funcData  OnTableGetData
	funcEvent OnTableEvent
	funcUrl   OnTableUrl
}

func NewTable(id string, header []string, gd OnTableGetData) *HTable {
	h := make([]string, len(header), len(header))
	copy(h, header)
	t := &HTable{newElem(id, "table", HtmlTable), h, make([]int, len(header), len(header)), false, false, gd, nil, nil}
	GlobalTables[id] = t
	t.self = t
	return t
}

func NewToolTable(id string, search bool, header []string, gd OnTableGetData, event OnTableEvent) *HTable {
	t := NewTable(id, header, gd)
	t.Tool = true
	t.Search = search
	t.funcEvent = event
	return t
}

func (table *HTable) SetColumnType(index int, c TableColumnType) {
	table.ColType[index] = int(c)
}
func (table *HTable) SetUrlHandler(funcUrl OnTableUrl) {
	table.funcUrl = funcUrl
}

var TempTable = `
<script id="__TABLEID___templet" type="text/html">

  <div class="layui-row layui-col-xs12 layui-col-sm12 layui-col-md12 layui-col-space3">
  {{# layui.each(Object.keys(d.data), function(index, item){ }}
	<div class="layui-col-xs3 layui-col-sm3 layui-col-md3">
	<div class="layui-form-item">
	<label class="layui-text">&nbsp;&nbsp;{{ item }}</label>
	</div>
	</div>
	<div class="layui-col-xs9 layui-col-sm9 layui-col-md9">
	<div class="layui-form-item">
	<input type="text" id="__TABLEID___col{{ index }}" class="layui-input" value="{{ d.data[item] }}">
	</div>
	</div>
  {{# }); }}
  {{#  if(d.oper != "detail"){ }}
	<div class="layui-col-xs4 layui-col-sm4 layui-col-md4 layui-col-xs-offset5 layui-col-sm-offset5 layui-col-md-offset5">
    <button class="layui-btn" onclick="__TABLEID___edit('{{ d.oper }}')">确认</button>
	</div>
  {{#  } }}
  </div>

</script>

`
var HtmlTable = `
<script type="text/html" id="{{.Id}}_toolbarHeader">
  <div class="{{if not .Search}}layui-hide{{end}} layui-container">
	<div class="layui-inline">
    <input class="layui-input layui-input-sm" id="{{.Id}}_search" autocomplete="off">
  	</div>
    <button class="layui-btn layui-btn-sm" lay-event="search">Search</button>
  </div>
</script>

<script type="text/html" id="{{.Id}}_toolbar">
  <a class="layui-btn layui-btn-primary layui-btn-xs" lay-event="detail">查看</a>
  <a class="layui-btn layui-btn-xs" lay-event="edit">编辑</a>
  <a class="layui-btn layui-btn-danger layui-btn-xs" lay-event="del">删除</a>
</script>

<table class="layui-hide" id="{{.Id}}" lay-filter="{{.Id}}"></table>

	<script>
	{{$TId:=.Id}}
	function {{RawString .Id}}_edit(op) {
		var $ = layui.jquery;
		var table = layui.table;
		var url = "/api/table?event_id={{.Id}}&oper="+op{{range $k,$v:=.Header}}+"&{{$k}}="+$("#{{$TId}}_col{{$k}}").val(){{end}};
		$.get(url,function(res){
			var r = JSON.parse(res);
			ret = r.code;
            if(ret == 0){
                layer.msg('操作成功');
        		layer.closeAll('page');
				table.reload('{{.Id}}');
            }else{
                layer.msg(r.msg);
            }
	    });
	}

	
	layui.use(['jquery', 'table', 'laytpl'], function () {
	        var $ = layui.jquery,
	            table = layui.table,
				laytpl = layui.laytpl;

			table.render({
			    elem: '#{{.Id}}'
			    ,url:'/api/table?event_id={{.Id}}&oper=data'
				{{if .Tool}},toolbar:'#{{.Id}}_toolbarHeader'{{end}}
				{{if .Tool}},defaultToolbar: [{title: '新加',layEvent: 'LAYTABLE_ADD',icon: 'layui-icon-addition'},'filter', 'exports', 'print']{{end}}
				,page:true
				,limit: 30
				,limits:[10,30,50,70,90,100]
				,id:'{{.Id}}'
			    ,cols: [[
	{{range $i,$v:=.Header}} {{if gt $i 0}},{{end}} {field:'col{{$i}}', title: '{{$v}}',align: 'center', {{$ty:=IntSliceElem $.ColType $i}}
		{{if eq $i 0}}sort: true,{{end}}
		{{if eq $ty 1}} event: 'img_col{{$i}}', templet: function(d){return '<a href="javascript:;"><img src='+d.col{{RawInt $i}}+'></a>'} {{end}}
		{{if eq $ty 2}} event: 'url_col{{$i}}', templet: function(d){return '<a class="layui-table-link" href="javascript:;">'+d.col{{RawInt $i}}+'</a>'} {{end}} }
	{{end}}
	{{if .Tool}},{fixed: 'right', title:'操作', toolbar: '#{{.Id}}_toolbar', width:163}{{end}}
			    ]]
			  });

			//头工具栏事件
		  table.on('toolbar({{.Id}})', function(obj){
		    switch(obj.event){
		      case 'search':
		          var cont = $('#{{.Id}}_search');
			      //执行重载
			      table.reload('{{.Id}}', {
			        page: {
			          curr: 1 //重新从第 1 页开始
			        }
			        ,where: {
			          key: {
			            search: cont.val()
			          }
			        }
			      }, 'data');
		      break;
		      //自定义头工具栏右侧图标 - 提示
		      case 'LAYTABLE_ADD':
				var td = { //数据
				  "oper":"add"
				  ,"data":{ {{range $i,$v:=.Header}} {{if gt $i 0}},{{end}}{{$v}}:""{{end}}}
				}
				
				var getTpl = document.getElementById('{{.Id}}_templet').innerHTML;
				laytpl(getTpl).render(td, function(html){
					layer.open({
					  	type: 1,
						title: "添加",
						area: ['50%', '60%'],
					  	content: html
					});	
				});
		      break;
		    };
		  });

		  //监听工具条
		  table.on('tool({{.Id}})', function(obj){
			var data = Object.values(obj.data);
		    if(obj.event === 'detail'){
			  var td = { //数据
				  "oper":"detail"
				  ,"data":{ {{range $i,$v:=.Header}} {{if gt $i 0}},{{end}}{{$v}}:data[{{$i}}]{{end}}}
				}
				var getTpl = document.getElementById('{{.Id}}_templet').innerHTML;
				laytpl(getTpl).render(td, function(html){
					layer.open({
					  	type: 1,
						title: "查看",
						area: ['50%', '60%'],
					  	content: html
					});	
				});
		    } else if(obj.event === 'del'){
		      layer.confirm('确认删除', function(index){
				$.get("/api/table?event_id={{.Id}}&oper=del"{{range $i,$v:=.Header}}+"&{{$i}}="+data[{{$i}}]{{end}},function(res){
					var r = JSON.parse(res);
					ret = r.code;
	                if(ret == 0){
		        		obj.del();
	                    layer.msg('删除成功');
		        		layer.close(index);
	                }else{
	                    layer.msg(r.msg);
	                }
            	});
		      });
		    } else if(obj.event === 'edit'){
		      var td = { //数据
				  "oper":"edit"
				  ,"data":{ {{range $i,$v:=.Header}} {{if gt $i 0}},{{end}}{{$v}}:data[{{$i}}]{{end}}}
				}
				var getTpl = document.getElementById('{{.Id}}_templet').innerHTML;
				laytpl(getTpl).render(td, function(html){
					layer.open({
					  	type: 1,
						title: "编辑",
						area: ['50%', '60%'],
					  	content: html
					});	
				});
		    }else{
					var event=obj.event;
					if(event.indexOf("img_")==0)
					{
						var img="";
						{{range $i,$v:=.Header}}
						if(event === 'img_col{{$i}}')
						{
							img=obj.data.col{{RawInt $i}}
						}
						{{end}}
						layer.open({
					        type: 1,
							title: "图片",
					        //skin: 'layui-layer-rim', //加上边框
					        area: ['30%', '50%'], //宽高
					        shadeClose: true, //开启遮罩关闭
					        end: function (index, layero) {
					            return false;
					        },
					        content: '<div style="text-align:center"><img src="' +img + '" /></div>'
					    });
					}else if(event.indexOf("url_")==0)
					{
						var url="";
						{{range $i,$v:=.Header}}
						if(event === 'url_col{{$i}}')
						{
							url=obj.data.col{{RawInt $i}}
						}
						{{end}}
						layer.open({
						  	type: 2,
							title: "预览",
							area: ['50%', '70%'], //宽高
						  	content: "/api/table?event_id={{.Id}}&oper=url&href="+url
						});
					}
				}
		  });

	});
	</script>

`

/*
	<style type="text/css">
	.layui-table-cell{
		text-align:center;
		height: 100%;
        max-width: 100%;
	}
	</style>
*/

func (table *HTable) Clone() HtmlElem {
	nt := &HTable{}
	*nt = *table
	h := make([]string, len(table.Header), len(table.Header))
	copy(h, table.Header)
	c := make([]int, len(table.Header), len(table.Header))
	copy(c, table.ColType)
	nt.Header = h
	nt.ColType = c
	nt.ElemBase.clone(table.ElemBase)
	return nt
}
func (table *HTable) Render() string {
	return strings.Replace(TempTable, "__TABLEID__", table.ID(), -1) + table.ElemBase.Render()
}
