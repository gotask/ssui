// app.go
package ssui

import (
	"net/http"
	"strings"
	"time"
)

type HAppData struct {
	Create time.Time
	Frames map[string]*Frame
}

type HApp struct {
	Address string
	Handler *http.ServeMux
	Global  *HAppData
	Data    map[string]*HAppData
}

func NewApp(addr string) *HApp {
	return &HApp{addr, &http.ServeMux{}, &HAppData{time.Now(), make(map[string]*Frame, 0)}, make(map[string]*HAppData, 0)}
}

func (a *HApp) AddFrame(f *Frame) *HApp {
	a.Global.Frames[f.Router] = f
	return a
}

func (a *HApp) Reset(token string) {
	delete(a.Data, token)
}

func (a *HApp) GetFrame(token, router string) *Frame {
	ha, ok := a.Data[token]
	if !ok {
		ha = &HAppData{time.Now(), make(map[string]*Frame, 0)}
		a.Data[token] = ha
	}
	if f, ok := ha.Frames[router]; ok {
		return f
	}
	if f, ok := a.Global.Frames[router]; ok {
		nf := f.Clone().(*Frame)
		ha.Frames[router] = nf
		return nf
	}
	return nil
}

func (a *HApp) GetElem(token, router, id string) HtmlElem {
	f := a.GetFrame(token, router)
	if f != nil {
		if e, o := f.Events[id]; o {
			return e
		}
	}
	return nil
}

func (a *HApp) Run() error {
	h := a.Handler
	//css js
	h.Handle("/layui/", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".css") {
			resp.Header().Set("content-type", "text/css; charset=utf-8")
		}
		http.StripPrefix("/layui/", http.FileServer(Layui)).ServeHTTP(resp, req)
	}))

	for r, _ := range a.Global.Frames {
		tempr := r
		h.HandleFunc(r, func(w http.ResponseWriter, r *http.Request) {
			/*if r.URL.Path != tempr {
				w.WriteHeader(404)
				return
			}*/
			token := GetToken(r)
			f := a.GetFrame(token, tempr)
			if f == nil {
				w.WriteHeader(404)
			} else {
				w.Write([]byte(f.Render(token)))
			}
		})
	}
	HandleButtonClick(a)
	HandleTableDel(a)
	HandleTableAdd(a)

	return http.ListenAndServe(a.Address, h)
}
