// checkbox.go
package ssui

import (
	"bytes"
	"html/template"
)

type HCheckBox struct {
	Id      string
	Text    string
	Checked bool
}

var HtmlCheckBox = `<div class="layui-form-item">
<input type="checkbox" lay-filter="layui_checkbox" id="{{.Id}}" title="{{.Text}}" {{if .Checked}}value="1" checked=""{{else}}value="0"{{end}}>
</div>`

func NewCheckBox(id, text string, checked bool) *HCheckBox {
	return &HCheckBox{id, text, checked}
}
func (b *HCheckBox) Type() string {
	return "checkbox"
}
func (b *HCheckBox) ID() string {
	return b.Id
}
func (b *HCheckBox) Clone() HtmlElem {
	return NewCheckBox(b.Id, b.Text, b.Checked)
}
func (b *HCheckBox) Render(token string) string {
	te := template.New("checkbox")
	t, e := te.Parse(HtmlCheckBox)
	if e != nil {
		return e.Error()
	}
	buf := bytes.NewBufferString("")
	e = t.Execute(buf, b)
	if e != nil {
		return e.Error()
	}
	return buf.String()
}
