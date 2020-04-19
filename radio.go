// radio.go
package ssui

type HRadio struct {
	*ElemBase
	SelIndex int //默认选择0
	Option   []string
}

var HtmlRadio = `<div class="layui-form-item {{if .Hide}}layui-hide{{end}}">
<div class="layui-inline">{{$ID:=.Id}}{{$SelIdx:=.SelIndex}}
<input type="text" class="layui-hide" id="{{$ID}}" value="{{$SelIdx}}">
{{range $i,$v:=.Option}}<input type="radio" name="{{$ID}}" lay-filter="{{$ID}}" value="{{$i}}" title="{{.}}" {{if eq $i $SelIdx}}checked=""{{end}}>{{end}}
</div>
<script>
	layui.use(['form'], function () {
	     var $ = layui.jquery,
	         form = layui.form;
	    //此处即为 radio 的监听事件
	    form.on('radio({{$ID}})', function(data){
	        $("#{{$ID}}").val(data.value)
	    });
	});
</script>
</div>`

func NewRadio(id string, selindex int, options []string) *HRadio {
	no := make([]string, len(options), len(options))
	copy(no, options)
	r := &HRadio{newElem(id, "radio", HtmlRadio), selindex, no}
	r.self = r
	return r
}
func (r *HRadio) Clone() HtmlElem {
	nr := NewRadio(r.Id, r.SelIndex, r.Option)
	nr.ElemBase.clone(r.ElemBase)
	return nr
}
