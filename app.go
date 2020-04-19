// app.go
package ssui

import (
	"fmt"
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
	Title   string
	Key     string //16 24 32
	Handler *http.ServeMux
	Global  *HAppData
	Data    map[string]*HAppData
	Group   []*PageGroup
	Home    *Frame
	Debug   bool
	Admin   bool
	LibPath string
}

func NewApp(addr string) *HApp {
	return &HApp{addr, "", "", &http.ServeMux{}, &HAppData{time.Now(), make(map[string]*Frame, 0)}, make(map[string]*HAppData, 0), nil, nil, false, false, ""}
}

func NewAdminApp(addr, title, key string) *HApp {
	defaultKey := "132353s4te7hsfg3"
	if len(key) < 16 {
		key = key + defaultKey
	}
	key = key[:16]
	return &HApp{addr, title, key, &http.ServeMux{}, &HAppData{time.Now(), make(map[string]*Frame, 0)}, make(map[string]*HAppData, 0), nil, nil, false, true, ""}
}

func NewDebugApp(addr, libpath string) *HApp {
	return &HApp{addr, "", "", &http.ServeMux{}, &HAppData{time.Now(), make(map[string]*Frame, 0)}, make(map[string]*HAppData, 0), nil, nil, true, false, libpath}
}

func NewAdminDebugApp(addr, title, key, libpath string) *HApp {
	defaultKey := "132353s4te7hsfg3"
	if len(key) < 16 {
		key = key + defaultKey
	}
	key = key[:16]
	return &HApp{addr, title, key, &http.ServeMux{}, &HAppData{time.Now(), make(map[string]*Frame, 0)}, make(map[string]*HAppData, 0), nil, nil, true, true, libpath}
}

func (a *HApp) UserValidCheck(user, router string) bool {
	if user == "admin" {
		return true
	}
	if router == "/authedit" {
		return false
	}
	return true
}

func (a *HApp) AddPageGroup(p *PageGroup) *HApp {
	a.Group = append(a.Group, p)
	return a
}

func (a *HApp) AddFrame(f *Frame) *HApp {
	if a.Home == nil {
		a.Home = f
	}
	a.Global.Frames[f.Router] = f
	return a
}

func (a *HApp) Reset(user string) {
	delete(a.Data, user)
}

func (a *HApp) ParseHttpParams(r *http.Request) map[string]string {
	r.ParseForm()
	params := make(map[string]string, 0)
	for k, v := range r.Form {
		params[k] = v[0]
	}

	c, e := r.Cookie("token")
	if e == nil {
		s, _ := Decrypt(c.Value, a.Key)
		params["token"] = s

		ss := strings.Split(s, "|")
		if len(ss) > 0 {
			params["username"] = ss[0]
		}
	}

	return params
}

func (a *HApp) getFrame(user, router string) *Frame {
	ha, ok := a.Data[user]
	if !ok {
		ha = &HAppData{time.Now(), make(map[string]*Frame, 0)}
		a.Data[user] = ha
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

func (a *HApp) GetFrame(user, router string) *Frame {
	f := a.getFrame(user, router)
	if f == nil {
		return f
	}
	if !a.UserValidCheck(user, router) {
		return nil
	}
	return f
}

func (a *HApp) GetElem(user, router, id string) HtmlElem {
	f := a.GetFrame(user, router)
	if f != nil {
		if e, o := f.Events[id]; o {
			return e
		}
	}
	return nil
}

func (a *HApp) addGroup(p *PageGroup) {
	for _, f := range p.Frames {
		a.AddFrame(f)
	}
	for _, g := range p.Groups {
		a.addGroup(g)
	}
}

// /mergely?fl=1&fr=2
func (a *HApp) SetMegelyFileFunc(f OnGetFile) {
	mergely.F = f
}

func (a *HApp) SetTokenExpireTime(d time.Duration) {
	Token_Expire_Time = d
}

func (a *HApp) Run() error {
	AuthEdit(a)
	for _, p := range a.Group {
		a.addGroup(p)
	}
	UserSetting(a)

	h := a.Handler
	//css js
	h.Handle("/uilib/", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".css") {
			resp.Header().Set("content-type", "text/css; charset=utf-8")
		}
		if a.Debug {
			UILib = http.Dir(a.LibPath)
		}
		http.StripPrefix("/uilib/", http.FileServer(UILib)).ServeHTTP(resp, req)
	}))

	h.Handle("/mergely", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		params := make(map[string]string, 0)
		for k, v := range r.Form {
			params[k] = v[0]
		}
		w.Write([]byte(mergely.Page(params["fl"], params["fr"])))
	}))

	for r, _ := range a.Global.Frames {
		tempr := r
		h.HandleFunc(r, func(w http.ResponseWriter, r *http.Request) {
			params := a.ParseHttpParams(r)
			user := params["username"]

			f := a.GetFrame(user, tempr)
			if f == nil {
				w.Write([]byte(Html404))
			} else {
				if f.OnLoad != nil {
					f.OnLoad(user)
				}
				fr := params["f"]
				if fr == "1" {
					w.Write([]byte(f.RenderFrame()))
				} else {
					w.Write([]byte(f.Render()))
				}
			}
		})
	}

	HandleButtonClick(a)
	HandleTable(a)
	HandleMergelay(a)

	if a.Admin {
		h.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			params := a.ParseHttpParams(r)
			user := params["username"]
			if user == "" {
				w.Write([]byte(LoginHtml))
				return
			}
			s := fmt.Sprintf(IndexHtml, a.Title, user)
			w.Write([]byte(s))
		}))
		h.Handle("/chpwd", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			params := a.ParseHttpParams(r)
			user := params["username"]
			if user == "" {
				w.Write([]byte(LoginHtml))
				return
			}
			w.Write([]byte(Chpwd_Html))
		}))

		HandleMenu(a)
		HandleClear(a)
		HandleLogin(a)
		HandleSignUp(a)
		HandleChpwd(a)
	}

	return http.ListenAndServe(a.Address, h)
}
