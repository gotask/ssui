// lineedit.go
package ssui

import (
	"fmt"
	"strings"
)

type HLineEdit struct {
	Id     string
	Prompt string
	Text   string
	Pass   bool
}

var HtmlLineEdit = `<div class="layui-form-item">
<input type="%s" id="%s" placeholder="%s" class="layui-input" value="%s">
</div>`

func NewLineEdit(id, prompt, text string, password bool) *HLineEdit {
	return &HLineEdit{id, prompt, text, password}
}
func (l *HLineEdit) Type() string {
	return "lineedit"
}
func (l *HLineEdit) ID() string {
	return l.Id
}
func (l *HLineEdit) Clone() HtmlElem {
	return NewLineEdit(l.Id, l.Prompt, l.Text, l.Pass)
}
func (l *HLineEdit) Render(token string) string {
	var buf strings.Builder
	ltype := "text"
	if l.Pass {
		ltype = "password"
	}
	fmt.Fprintf(&buf, HtmlLineEdit, ltype, l.Id, l.Prompt, l.Text)
	buf.WriteString("\n")
	return buf.String()
}
