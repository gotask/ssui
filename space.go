// space.go
package ssui

import (
	"strings"
)

type HSpace struct {
}

var HtmlSpace = `<div class="layui-form-item">
<label class="layui-form-label">&nbsp;</label>
</div>`

func NewSpace() *HSpace {
	return &HSpace{}
}
func (l *HSpace) Type() string {
	return "space"
}
func (l *HSpace) ID() string {
	return ""
}
func (l *HSpace) Clone() HtmlElem {
	return NewSpace()
}
func (l *HSpace) Render() string {
	var buf strings.Builder
	buf.WriteString(HtmlSpace)
	return buf.String()
}
