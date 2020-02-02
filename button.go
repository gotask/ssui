// button.go
package ssui

import (
	"fmt"
	"strings"
)

type OnButtonClick func(param map[string]string) *HResponse

type HButton struct {
	Id    string
	Text  string
	Event OnButtonClick
}

var HtmlButton = `<div class="layui-form-item">
<button type="button" class="layui-btn layui-btn-primary" id="%s" onclick="buttonClick('%s')">%s</button>
</div>`

func NewButton(id, text string, event OnButtonClick) *HButton {
	return &HButton{id, text, event}
}
func (b *HButton) Type() string {
	return "button"
}
func (b *HButton) ID() string {
	return b.Id
}
func (b *HButton) Clone() HtmlElem {
	return NewButton(b.Id, b.Text, b.Event)
}
func (b *HButton) Render(token string) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, HtmlButton, b.Id, b.Id, b.Text)
	buf.WriteString("\n")
	return buf.String()
}
