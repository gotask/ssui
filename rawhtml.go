// rawhtml
package ssui

import (
	"strings"
)

type HRawHtml struct {
	Text string
}

func NewRawHtml(text string) *HRawHtml {
	return &HRawHtml{text}
}
func (l *HRawHtml) Type() string {
	return "rawhtml"
}
func (l *HRawHtml) ID() string {
	return ""
}
func (l *HRawHtml) Clone() HtmlElem {
	return NewRawHtml(l.Text)
}
func (l *HRawHtml) Render() string {
	var buf strings.Builder
	buf.WriteString(l.Text)
	return buf.String()
}
