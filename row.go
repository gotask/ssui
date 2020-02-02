// row.go
package ssui

import (
	"strings"
)

type Row struct {
	Elems []HtmlElem
}

func NewRow() *Row {
	return &Row{make([]HtmlElem, 0)}
}
func (l *Row) AddElem(e HtmlElem) *Row {
	l.Elems = append(l.Elems, e)
	return l
}
func (l *Row) Type() string {
	return "row"
}
func (l *Row) ID() string {
	return ""
}
func (l *Row) Clone() HtmlElem {
	nl := NewRow()
	for _, r := range l.Elems {
		nl.AddElem(r.Clone())
	}
	return nl
}
func (l *Row) Render(token string) string {
	var buff strings.Builder
	buff.WriteString(`<div class="layui-form-item">`)

	for _, s := range l.Elems {
		buff.WriteString(`<div class="layui-inline">`)
		buff.WriteString(s.Render(token))
		buff.WriteString(`</div>`)
	}
	buff.WriteString(`</div>`)
	return buff.String()
}
