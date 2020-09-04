// frame.go
package ssui

import (
	"bytes"
	"html/template"
	"strings"
)

type OnFrameLoad func(user string)

type Frame struct {
	Router string
	Title  string
	Icon   string
	Elems  []HtmlElem
	Events map[string]HtmlElem
	OnLoad OnFrameLoad
}

//icon https://www.layui.com/doc/element/icon.html#table
func NewFrame(router, title, icon string, onload OnFrameLoad) *Frame {
	return &Frame{router, title, icon, make([]HtmlElem, 0), make(map[string]HtmlElem, 0), onload}
}
func (f *Frame) AddElem(e HtmlElem) *Frame {
	f.Elems = append(f.Elems, e)
	f.addevent(e)
	e.SetRouter(f.Router)
	return f
}
func (f *Frame) addevent(e HtmlElem) {
	if e.ID() != "" {
		f.Events[e.ID()] = e
	} else if e.Type() == "row" {
		for _, r := range e.(*Row).Elems {
			if r.ID() != "" {
				f.Events[r.ID()] = r
			} else {
				f.addevent(r)
			}
		}
	} else if e.Type() == "panel" {
		for _, r := range e.(*Panel).Elems {
			if r.elem.ID() != "" {
				f.Events[r.elem.ID()] = r.elem
			} else {
				f.addevent(r.elem)
			}
		}
	} else if e.Type() == "card" {
		f.addevent(e.(*HCard).Header)
		f.addevent(e.(*HCard).Body)
	}
}
func (f *Frame) Type() string {
	return "frame"
}
func (f *Frame) ID() string {
	return ""
}
func (f *Frame) Route() string {
	return f.Router
}
func (f *Frame) SetRouter(r string) {
	f.Router = r
}
func (f *Frame) SetValue(v string) {
	f.Title = v
}
func (f *Frame) Clone() HtmlElem {
	nf := NewFrame(f.Router, f.Title, f.Icon, f.OnLoad)
	for _, r := range f.Elems {
		nf.AddElem(r.Clone())
	}
	return nf
}

func (f *Frame) buildParams() string {
	//$("#date").val()
	param := "\"url_router=" + f.Router + "\""
	for _, s := range f.Events {
		if s.ID() == "" || s.Type() == "table" || s.Type() == "text" ||
			s.Type() == "button" || s.Type() == "echart" {
			continue
		}
		param += "+\"&"
		param += s.ID() + "=\"+" + "$(\"#" + s.ID() + "\").val()"
	}
	param += ";\n}\n"
	return param + f.buildFunc()
}

func (f *Frame) buildFunc() string {
	fun := "function getRouter(){return \""
	fun += f.Router
	fun += "\"}\n"
	return fun
}

func (f *Frame) RenderFrame() string {
	var buff strings.Builder

	buff.WriteString(HtmlHeader)

	buff.WriteString(HtmlScript)
	buff.WriteString("return " + f.buildParams())
	buff.WriteString(HtmlScriptFrame)

	for _, s := range f.Elems {
		buff.WriteString(s.Render())
	}

	buff.WriteString(HtmlFooter)
	return buff.String()
}

func (f *Frame) Render() string {
	var buff strings.Builder

	te := template.New("frame")
	t, e := te.Parse(HtmlPage1)
	if e != nil {
		return e.Error()
	}
	b := bytes.NewBufferString("")
	e = t.Execute(b, f)
	if e != nil {
		return e.Error()
	}
	buff.WriteString(b.String())

	buff.WriteString(HtmlHeader)

	buff.WriteString(HtmlScript)
	buff.WriteString("return " + f.buildParams())
	buff.WriteString(HtmlScriptPage)

	for _, s := range f.Elems {
		buff.WriteString(s.Render())
	}

	buff.WriteString(HtmlFooter)

	buff.WriteString(HtmlPage2)

	return buff.String()
}
