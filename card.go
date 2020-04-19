// card
package ssui

type HCard struct {
	*ElemBase
	Header HtmlElem
	Body   HtmlElem
}

var HtmlCard = `<div class="layui-card {{if .Hide}}layui-hide{{end}}">
  <div class="layui-card-header">{{Render .Header}}</div>
  <div class="layui-card-body">{{Render .Body}}</div>
</div>`

func NewCard(header, body HtmlElem) *HCard {
	c := &HCard{newElem("", "card", HtmlCard), header, body}
	c.self = c
	return c
}
func (c *HCard) Clone() HtmlElem {
	nc := NewCard(c.Header, c.Body)
	nc.ElemBase.clone(c.ElemBase)
	return nc
}
