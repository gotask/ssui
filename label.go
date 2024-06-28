// label.go
package ssui

type HLabel struct {
	*ElemBase
	Text string
}

// layui-form-item
var HtmlLabel = `<div class="layui-form-item {{if .Hide}}layui-hide{{end}}">
<label class="layui-form-label" style="text-align:{{.Align}};">{{RawString .Text}}</label>
</div>`

func NewLabel(text string) *HLabel {
	l := &HLabel{newElem("", "label", HtmlLabel), text}
	l.Align = "left"
	l.self = l
	return l
}

func (l *HLabel) Clone() HtmlElem {
	nl := NewLabel(l.Text)
	nl.ElemBase.clone(l.ElemBase)
	return nl
}
