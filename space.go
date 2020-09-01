// space.go
package ssui

import (
	"strings"
)

type HSpace struct {
	*ElemBase
}

var HtmlSpace = `<div class="layui-form-item">
<label class="layui-form-label">&nbsp;</label>
</div>`

func NewSpace() *HSpace {
	return &HSpace{&ElemBase{}}
}
func (l *HSpace) Type() string {
	return "space"
}
func (l *HSpace) ID() string {
	return ""
}
func (l *HSpace) Clone() HtmlElem {
	nl := NewSpace()
	nl.ElemBase.clone(l.ElemBase)
	return nl
}
func (l *HSpace) Render() string {
	var buf strings.Builder
	buf.WriteString(HtmlSpace)
	return buf.String()
}
