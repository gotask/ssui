// radio.go
// select.go
package ssui

import (
	"bytes"
	"html/template"
)

type HRadio struct {
	Id       string
	SelIndex int //默认选择0
	Option   []string
}

var HtmlRadio = `<div class="layui-form-item">
<div class="layui-input-block">{{$ID:=.Id}}{{$SelIdx:=.SelIndex}}
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
	return &HRadio{id, selindex, no}
}
func (s *HRadio) Type() string {
	return "radio"
}
func (s *HRadio) ID() string {
	return s.Id
}
func (s *HRadio) Clone() HtmlElem {
	return NewRadio(s.Id, s.SelIndex, s.Option)
}
func (s *HRadio) Render(token string) string {
	te := template.New("radio")
	t, e := te.Parse(HtmlRadio)
	if e != nil {
		return e.Error()
	}
	buf := bytes.NewBufferString("")
	e = t.Execute(buf, s)
	if e != nil {
		return e.Error()
	}
	return buf.String()
}
