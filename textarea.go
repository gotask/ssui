// textarea.go
package ssui

import (
	"fmt"
	"strings"
)

type HTextArea struct {
	Id     string
	Prompt string
	Text   string
}

var HtmlTextArea = `<div class="layui-form-item">
<textarea id="%s" placeholder="%s" class="layui-textarea">%s</textarea>
</div>`

func NewTextArea(id, prompt, text string) *HTextArea {
	return &HTextArea{id, prompt, text}
}
func (l *HTextArea) Type() string {
	return "textarea"
}
func (l *HTextArea) ID() string {
	return l.Id
}
func (l *HTextArea) Clone() HtmlElem {
	return NewTextArea(l.Id, l.Prompt, l.Text)
}
func (l *HTextArea) Render(token string) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, HtmlTextArea, l.Id, l.Prompt, l.Text)
	buf.WriteString("\n")
	return buf.String()
}
