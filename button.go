// button.go
package ssui

// OnButtonClick username 当前用户 param 当前页面所有可写组件的值集合
type OnButtonClick func(username string, param map[string]string) *HResponse

type HButton struct {
	*ElemBase
	Text  string
	Event OnButtonClick
}

// layui-form-item
var HtmlButton = `<div class="layui-form-item {{if .Hide}}layui-hide{{end}}" style="text-align:{{.Align}};">
<button type="button" class="layui-btn layui-btn-primary {{if .Disable}}layui-disabled{{end}}" id="{{.Id}}" {{if not .Disable}}onclick="buttonClick('{{.Id}}')"{{end}}>{{RawString .Text}}</button>
</div>`

// NewButton OnButtonClick 可用app.GetElem来获取组件并修改其值
func NewButton(id, text string, event OnButtonClick) *HButton {
	b := &HButton{newElem(id, "button", HtmlButton), text, event}
	b.self = b
	return b
}
func (b *HButton) Clone() HtmlElem {
	nb := NewButton(b.Id, b.Text, b.Event)
	nb.ElemBase.clone(b.ElemBase)
	return nb
}
