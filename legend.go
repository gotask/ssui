// legend.go
package ssui

import (
	"fmt"
	"strings"
)

type HLegend struct {
	Text string
}

var HtmlLegend = `<fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
  <legend>%s</legend>
</fieldset>`

func NewLegend(text string) *HLegend {
	return &HLegend{text}
}
func (l *HLegend) Type() string {
	return "legend"
}
func (l *HLegend) ID() string {
	return ""
}
func (l *HLegend) Clone() HtmlElem {
	return NewLegend(l.Text)
}
func (l *HLegend) Render(token string) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, HtmlLegend, l.Text)
	buf.WriteString("\n")
	return buf.String()
}
