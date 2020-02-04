// label.go
package ssui

import (
	"fmt"
	"strings"
)

type HLabel struct {
	Text string
}

var HtmlLabel = `<div class="layui-form-item">
<label>%s</label>
</div>`

func NewLabel(text string) *HLabel {
	return &HLabel{text}
}
func (l *HLabel) Type() string {
	return "label"
}
func (l *HLabel) ID() string {
	return ""
}
func (l *HLabel) Clone() HtmlElem {
	return NewLabel(l.Text)
}
func (l *HLabel) Render(token string) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, HtmlLabel, l.Text)
	buf.WriteString("\n")
	return buf.String()
}
