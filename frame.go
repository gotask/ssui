// frame.go
package ssui

import (
	"bytes"
	"html/template"
	"strings"
)

type HtmlElem interface {
	ID() string
	Type() string
	Clone() HtmlElem
	Render(token string) string
}

type Frame struct {
	Router string
	Title  string
	Header string
	Footer string
	Elems  []HtmlElem
	Events map[string]HtmlElem
}

func NewFrame(router, title, header, footer string) *Frame {
	return &Frame{router, title, header, footer, make([]HtmlElem, 0), make(map[string]HtmlElem, 0)}
}
func (f *Frame) AddElem(e HtmlElem) *Frame {
	f.Elems = append(f.Elems, e)
	if e.ID() != "" {
		f.Events[e.ID()] = e
	}
	if e.Type() == "row" {
		for _, r := range e.(*Row).Elems {
			if r.ID() != "" {
				f.Events[r.ID()] = r
			}
		}
	}
	return f
}
func (f *Frame) Type() string {
	return "frame"
}
func (f *Frame) ID() string {
	return ""
}
func (f *Frame) Clone() HtmlElem {
	nf := NewFrame(f.Router, f.Title, f.Header, f.Footer)
	for _, r := range f.Elems {
		nf.AddElem(r.Clone())
	}
	return nf
}
func (f *Frame) Render(token string) string {
	var buff strings.Builder
	buff.WriteString(HtmlHeader)

	for _, s := range f.Elems {
		buff.WriteString(s.Render(token))
	}
	//$("#date").val()
	param := ""
	cnt := 0
	for _, s := range f.Events {
		if s.ID() == "" || s.Type() == "table" ||
			s.Type() == "row" || s.Type() == "button" {
			continue
		}
		if cnt != 0 {
			param += "+"
		}
		param += "\""
		if cnt != 0 {
			param += "&"
		}
		param += s.ID() + "=\"+" + "$(\"#" + s.ID() + "\").val()"
		cnt++
	}
	if cnt == 0 {
		param = `""`
	}
	param += "}\n"

	buff.WriteString(HtmlScript1)
	buff.WriteString("return " + param)
	buff.WriteString(HtmlScript2)
	buff.WriteString(HtmlFooter)
	te := template.New("frame")
	t, e := te.Parse(buff.String())
	if e != nil {
		return e.Error()
	}
	b := bytes.NewBufferString("")
	e = t.Execute(b, f)
	if e != nil {
		return e.Error()
	}
	return b.String()
}
