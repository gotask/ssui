// select.go
package ssui

import (
	"bytes"
	"html/template"
)

type HSelect struct {
	Id       string
	SelIndex int
	Option   []string
}

var HtmlSelect = `<div class="layui-form-item">
<select id="{{.Id}}">
{{range .Option}}<option value="{{.}}">{{.}}</option>{{end}}
</select>
<script>
	layui.use(['form'], function () {
	     var $ = layui.jquery,
	         form = layui.form;
	    
		 $("#{{.Id}}").val('{{SliceElem .Option .SelIndex}}');
         form.render('select');
	});
</script>
</div>`

func NewSelect(id string, selindex int, options []string) *HSelect {
	no := make([]string, len(options), len(options))
	copy(no, options)
	return &HSelect{id, selindex, no}
}
func (s *HSelect) Type() string {
	return "select"
}
func (s *HSelect) ID() string {
	return s.Id
}
func (s *HSelect) Clone() HtmlElem {
	return NewSelect(s.Id, s.SelIndex, s.Option)
}
func (s *HSelect) Render(token string) string {
	te := template.New("select")
	te = te.Funcs(template.FuncMap{"SliceElem": SliceElem})
	t, e := te.Parse(HtmlSelect)
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
