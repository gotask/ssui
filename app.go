// app.go
package ssui

import (
	"fmt"
	"github.com/gotask/gost/stnet"
	"net/http"
	"strings"
	"time"
)

var (
	Admin_User_Name   = "adminxyz"           //超级管理员账户，要自己先注册
	Token_Expire_Time = 365 * 24 * time.Hour //默认一年
)

type UserAuthCheck func(user, router string) bool

type HAppData struct {
	Create time.Time
	Frames map[string]*Frame
}

type HApp struct {
	address     string               //监听地址
	title       string               //标题
	key         string               //秘钥，用于token加密，密码加密，长度16 24 32
	handler     *stnet.HttpHandler   //url handler(http.ServeMux)
	global      *HAppData            //全局数据
	data        map[string]*HAppData //用户数据
	group       []*PageGroup         //menu
	home        *Frame               //主页
	debug       bool                 //debug模式下使用uilib目录
	admin       bool                 //admin模式
	libPath     string               //uilib路径
	authCheck   UserAuthCheck        //验证用户访问权限
	openRegiste bool                 //是否开放注册
}

func NewApp(addr string) *HApp {
	return &HApp{address: addr,
		handler: &stnet.HttpHandler{}, //&http.ServeMux{},
		global:  &HAppData{time.Now(), make(map[string]*Frame, 0)},
		data:    make(map[string]*HAppData, 0),
	}
}

func NewAdminApp(addr, title, key string) *HApp {
	defaultKey := "132353s4te7hsfg3"
	if len(key) < 16 {
		key = key + defaultKey
	}
	key = key[:16]
	return &HApp{address: addr,
		title:   title,
		key:     key,
		handler: &stnet.HttpHandler{}, //&http.ServeMux{},
		global:  &HAppData{time.Now(), make(map[string]*Frame, 0)},
		data:    make(map[string]*HAppData, 0),
		admin:   true,
	}
}

func NewDebugApp(addr, libpath string) *HApp {
	return &HApp{address: addr,
		handler: &stnet.HttpHandler{}, //&http.ServeMux{},
		global:  &HAppData{time.Now(), make(map[string]*Frame, 0)},
		data:    make(map[string]*HAppData, 0),
		debug:   true,
		libPath: libpath,
	}
}

func NewAdminDebugApp(addr, title, key, libpath string) *HApp {
	defaultKey := "132353s4te7hsfg3"
	if len(key) < 16 {
		key = key + defaultKey
	}
	key = key[:16]
	return &HApp{address: addr,
		title:   title,
		key:     key,
		handler: &stnet.HttpHandler{}, //&http.ServeMux{},
		global:  &HAppData{time.Now(), make(map[string]*Frame, 0)},
		data:    make(map[string]*HAppData, 0),
		admin:   true,
		debug:   true,
		libPath: libpath,
	}
}

// 是否开放注册
func (a *HApp) OpenRegiste(r bool) {
	a.openRegiste = r
}

// 权限管理
func (a *HApp) SetAuthCheck(f UserAuthCheck) {
	a.authCheck = f
}

// token过期时间
func (a *HApp) SetTokenExpireTime(d time.Duration) {
	Token_Expire_Time = d
}

func (a *HApp) userValidCheck(user, router string) bool {
	if user == Admin_User_Name {
		return true
	}
	if router == "/authedit" {
		return false
	}
	if a.authCheck != nil {
		return a.authCheck(user, router)
	}
	if a.admin {
		return false
	}
	return true
}

// 添加组
func (a *HApp) AddPageGroup(p *PageGroup) *HApp {
	a.group = append(a.group, p)
	return a
}

// 添加页面
func (a *HApp) AddFrame(f *Frame) *HApp {
	if a.home == nil {
		a.home = f
	}
	a.global.Frames[f.Router] = f
	return a
}

// 重置用户缓存数据
func (a *HApp) Reset(user string) {
	delete(a.data, user)
}

// 获取http请求参数，从token中提取原始的username和token
func (a *HApp) ParseHttpParams(r *http.Request) map[string]string {
	r.ParseForm()
	params := make(map[string]string, 0)
	for k, v := range r.Form {
		params[k] = string(v[0])
	}

	c, e := r.Cookie("token")
	if e == nil {
		s, _ := Decrypt(c.Value, a.key)
		params["token"] = s

		ss := strings.Split(s, "|")
		if len(ss) > 0 {
			name := ss[0]
			usr := GetUser(a, name)
			if usr.Name == name {
				params["username"] = name
			}
		}
	}

	return params
}

func (a *HApp) getFrame(user, router string) *Frame {
	ha, ok := a.data[user]
	if !ok {
		ha = &HAppData{time.Now(), make(map[string]*Frame, 0)}
		a.data[user] = ha
	}
	if f, ok := ha.Frames[router]; ok {
		return f
	}
	if f, ok := a.global.Frames[router]; ok {
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
	if !a.userValidCheck(user, router) {
		return nil
	}
	return f
}

// GetElem 如需修改组件内容可通过此接口获取真实组件地址
func (a *HApp) GetElem(user, router, id string) HtmlElem {
	f := a.GetFrame(user, router)
	if f != nil {
		if e, o := f.Events[id]; o {
			return e
		}
	}
	return nil
}
func (a *HApp) GetElemWithVal(router, id string, param map[string]string) HtmlElem {
	user := param["username"]
	f := a.GetFrame(user, router)
	if f != nil {
		if e, o := f.Events[id]; o {
			if v, ok := param[id]; ok {
				e.SetValue(v)
			}
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

func (a *HApp) Run() error {
	AuthEdit(a)
	for _, p := range a.group {
		a.addGroup(p)
	}
	UserSetting(a)

	h := a.handler
	//css js
	h.Handle("/uilib/", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".css") {
			resp.Header().Set("content-type", "text/css; charset=utf-8")
		}
		if a.debug {
			UILib = http.Dir(a.libPath)
		}
		http.StripPrefix("/uilib/", http.FileServer(UILib)).ServeHTTP(resp, req)
	}))

	h.Handle("/mergely", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := a.ParseHttpParams(r)
		extend := "&event_id=" + params["event_id"] + "&url_router=" + params["url_router"]
		w.Write([]byte(MergelyPage(params["fl"]+extend, params["fr"]+extend)))
	}))

	for r, _ := range a.global.Frames {
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
	HandleUpload(a)

	if a.admin {
		h.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			params := a.ParseHttpParams(r)
			user := params["username"]
			if user == "" {
				w.Write([]byte(LoginHtml))
				return
			}
			s := fmt.Sprintf(IndexHtml, a.title, user)
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

	s := stnet.NewServer(10, 32)
	s.AddHttpService("http", a.address, 0, &HttpServer{}, h, 0)
	return s.Start()
	//return http.ListenAndServe(a.address, h)
}
