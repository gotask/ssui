// lineedit.go
package ssui

type HLineEdit struct {
	*ElemBase
	Prompt string
	Text   string
	Pass   bool
}

var HtmlLineEdit = `<div class="layui-form-item {{if .Hide}}layui-hide{{end}}">
<input {{if .Disable}}disabled=""{{end}} {{if .Pass}}type="password"{{else}}type="text"{{end}} id="{{.Id}}" placeholder="{{.Prompt}}" class="layui-input" value="{{.Text}}">
</div>`

func NewLineEdit(id, prompt, text string, password bool) *HLineEdit {
	l := &HLineEdit{newElem(id, "lineedit", HtmlLineEdit), prompt, text, password}
	l.self = l
	return l
}
func (l *HLineEdit) Clone() HtmlElem {
	nl := NewLineEdit(l.Id, l.Prompt, l.Text, l.Pass)
	nl.ElemBase.clone(l.ElemBase)
	return nl
}
func (l *HLineEdit) SetValue(v string) {
	l.Text = v
}
