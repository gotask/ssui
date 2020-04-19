// href.go
package ssui

type HHref struct {
	*ElemBase
	Text string
	URL  string
}

var HtmlHref = `<div class="layui-form-item {{if .Hide}}layui-hide{{end}}">
<a href="javascript:;" onclick="redirectUrl('{{.URL}}')" class="valid" style="text-decoration: underline;color: -webkit-link;">{{RawString .Text}}</a>
</div>`

func NewHref(text, url string) *HHref {
	h := &HHref{newElem("", "href", HtmlHref), text, url}
	h.self = h
	return h
}
func (h *HHref) Clone() HtmlElem {
	nh := NewHref(h.Text, h.URL)
	nh.ElemBase.clone(h.ElemBase)
	return nh
}
