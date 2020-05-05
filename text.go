// text
package ssui

import (
	"fmt"
	"strings"
)

type HText struct {
	*ElemBase
	Text string
}

func (t *HText) SetText(s string) {
	t.Text = s
}

/*var HtmlText = `<div class="layui-form-item %s">
<xmp class="layui-code">%s</xmp>
</div>
`
*/
var HtmlText = `<div class="layui-form-item %s">
<pre class="layui-code">%s</pre>
</div>`

func NewText(id, text string) *HText {
	t := &HText{newElem(id, "text", HtmlText), text}
	t.self = t
	return t
}
func (t *HText) Clone() HtmlElem {
	nt := NewText(t.Id, t.Text)
	nt.ElemBase.clone(t.ElemBase)
	return nt
}
func (t *HText) Render() string {
	var buff strings.Builder
	hide := ""
	if t.Hide {
		hide = "layui-hide"
	}
	fmt.Fprintf(&buff, HtmlText, hide, t.Text)
	return buff.String()
}
