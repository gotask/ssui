// textarea.go
package ssui

type HTextArea struct {
	*ElemBase
	Prompt string
	Text   string
	Rows   int
}

var HtmlTextArea = `<div class="layui-form-item {{if .Hide}}layui-hide{{end}}">
<textarea id="{{.Id}}" {{if .Disable}}disabled=""{{end}} placeholder="{{.Prompt}}" class="layui-textarea" rows="{{.Rows}}">{{RawString .Text}}</textarea>
</div>`

func NewTextArea(id, prompt, text string) *HTextArea {
	t := &HTextArea{newElem(id, "textarea", HtmlTextArea), prompt, text, 17}
	t.self = t
	return t
}
func (t *HTextArea) Clone() HtmlElem {
	nt := NewTextArea(t.Id, t.Prompt, t.Text)
	nt.Rows = t.Rows
	nt.ElemBase.clone(t.ElemBase)
	return nt
}
func (t *HTextArea) SetValue(v string) {
	t.Text = v
}
