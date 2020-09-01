// rawhtml
package ssui

import (
	"strings"
)

type HRawHtml struct {
	*ElemBase
	Text string
}

func NewRawHtml(text string) *HRawHtml {
	return &HRawHtml{&ElemBase{}, text}
}
func (l *HRawHtml) Type() string {
	return "rawhtml"
}
func (l *HRawHtml) ID() string {
	return ""
}
func (l *HRawHtml) Route() string {
	return ""
}
func (l *HRawHtml) Clone() HtmlElem {
	nl := NewRawHtml(l.Text)
	nl.ElemBase.clone(l.ElemBase)
	return nl
}
func (l *HRawHtml) Render() string {
	var buf strings.Builder
	buf.WriteString(l.Text)
	return buf.String()
}
