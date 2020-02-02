// timepicker.go
package ssui

import (
	"bytes"
	"html/template"
)

type HTimePicker struct {
	Id string
	//y代表年M代表月,以此类推,例如: yyyy-MM-dd HH:mm:ss yyyy年M月 yyyy年的M月某天晚上，大概H点 dd/MM/yyyy ...
	Format string
	Value  int64 //1970年以来的ms数
}

var HtmlTimePicker = `<div class="layui-form-item">
<input type="text" class="layui-input" id="{{.Id}}">
<script>
	layui.use(['form','laydate'], function () {
	     var $ = layui.jquery,
	         form = layui.form,
			 laydate = layui.laydate;

	    //日期
        laydate.render({
            elem: '#{{.Id}}'
			,type: 'datetime'
            ,format: '{{.Format}}'
            ,value: new Date({{if gt .Value 0}}{{.Value}}{{end}})
            ,isInitValue: true
        });
	});
</script>
</div>`

func NewTimePicker(id, format string, val int64) *HTimePicker {
	return &HTimePicker{id, format, val}
}
func (s *HTimePicker) Type() string {
	return "timepicker"
}
func (s *HTimePicker) ID() string {
	return s.Id
}
func (s *HTimePicker) Clone() HtmlElem {
	return NewTimePicker(s.Id, s.Format, s.Value)
}
func (s *HTimePicker) Render(token string) string {
	te := template.New("timepicker")
	t, e := te.Parse(HtmlTimePicker)
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
