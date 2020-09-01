// row.go
package ssui

import (
	"fmt"
	"strings"
)

type PanelElem struct {
	phoneWidth int
	deskWidth  int
	elem       HtmlElem
}

type Panel struct {
	*ElemBase
	phoneWidth int
	deskWidth  int
	margin     int
	Elems      []PanelElem
}

//width value:1-12(percent) margin 1-30(px)
func NewPanel(phoneWidth, deskWidth, margin int) *Panel {
	if phoneWidth < 1 || phoneWidth > 12 {
		phoneWidth = 12
	}
	if deskWidth < 1 || deskWidth > 12 {
		deskWidth = 12
	}
	if margin < 1 || margin > 30 {
		margin = 0
	}
	return &Panel{&ElemBase{}, phoneWidth, deskWidth, margin, make([]PanelElem, 0)}
}

//width value:1-12
func (l *Panel) AddElem(e HtmlElem, phoneWidth, deskWidth int) *Panel {
	if phoneWidth < 1 || phoneWidth > 12 {
		phoneWidth = 12
	}
	if deskWidth < 1 || deskWidth > 12 {
		deskWidth = 12
	}
	l.Elems = append(l.Elems, PanelElem{phoneWidth, deskWidth, e})
	return l
}
func (l *Panel) Type() string {
	return "panel"
}
func (l *Panel) ID() string {
	return ""
}
func (l *Panel) SetRouter(r string) {
	l.ElemBase.SetRouter(r)
	for _, v := range l.Elems {
		v.elem.SetRouter(r)
	}
}
func (l *Panel) Clone() HtmlElem {
	nl := NewPanel(l.phoneWidth, l.deskWidth, l.margin)
	for _, r := range l.Elems {
		nl.AddElem(r.elem.Clone(), r.phoneWidth, r.deskWidth)
	}
	nl.ElemBase.clone(l.ElemBase)
	return nl
}
func (l *Panel) Render() string {
	var buff strings.Builder
	//buff.WriteString(`<div class="layui-row">`)
	fmt.Fprintf(&buff, "<div class=\"layui-row layui-col-xs%d layui-col-sm%d layui-col-md%d layui-col-space%d\">", l.phoneWidth, l.phoneWidth, l.deskWidth, l.margin)
	if len(l.Elems) == 0 {
		buff.WriteString("<label>&nbsp;</label>")
	}
	for _, s := range l.Elems {
		fmt.Fprintf(&buff, "<div class=\"layui-col-xs%d layui-col-sm%d layui-col-md%d\">", s.phoneWidth, s.phoneWidth, s.deskWidth)
		buff.WriteString(s.elem.Render())
		buff.WriteString(`</div>`)
	}
	buff.WriteString(`</div>`)
	return buff.String()
}
