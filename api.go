// event.go
package ssui

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**************************************************************************
			event_id 事件发起者ID
			url_router 事件页面路由
			其他key val分别表示页面控件ID和Val
**************************************************************************/
func HandleButtonClick(a *HApp) {
	a.handler.HandleFunc("/button_click", func(w http.ResponseWriter, r *http.Request) {
		params := a.ParseHttpParams(r)
		event_id, ok := params["event_id"]
		url_router, ok1 := params["url_router"]
		user := params["username"]
		res := &HResponse{Error: "error param"}
		for {
			if !ok || !ok1 {
				break
			}
			e := a.GetElem(user, url_router, event_id)
			if e == nil {
				break
			}
			b := e.(*HButton)
			if b.Event == nil {
				break
			}
			res = b.Event(user, params)
			break
		}

		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}

func addMenu(a *HApp, p *PageGroup, m *MenuChild, user string) {
	m.Href = ""
	m.Icon = p.Icon
	m.Title = p.Title
	m.Target = "_self"

	for _, f := range p.Frames {
		if a.userValidCheck(user, f.Router) {
			m.Child = append(m.Child, MenuChild{nil, f.Router, f.Icon, "_self", f.Title})
		}
	}
	for _, g := range p.Groups {
		nm := MenuChild{}
		addMenu(a, g, &nm, user)
		m.Child = append(m.Child, nm)
	}
}
func HandleMenu(a *HApp) {
	a.handler.HandleFunc("/api/init", func(w http.ResponseWriter, r *http.Request) {
		params := a.ParseHttpParams(r)
		res := &Menu{}
		user := params["username"]

		if a.home == nil {
			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
			return
		}

		res.HomeInfo.Title = a.title
		res.HomeInfo.Href = a.home.Router
		res.LogoInfo.Title = a.title
		res.LogoInfo.Image = "/uilib/images/logo.png"
		res.LogoInfo.Href = a.home.Router

		for _, p := range a.group {
			if p.Title == "AuthEdit" && user != Admin_User_Name {
				continue
			}
			nm := MenuChild{}
			addMenu(a, p, &nm, user)
			res.MenuInfo = append(res.MenuInfo, nm)
		}
		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}

func HandleClear(a *HApp) {
	a.handler.HandleFunc("/api/clear", func(w http.ResponseWriter, r *http.Request) {
		params := a.ParseHttpParams(r)
		user := params["username"]
		if user == "" {
			res := &ApiRsp{1, "清理缓存失败"}
			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
			return
		}
		a.Reset(user)
		res := &ApiRsp{0, "清理缓存成功"}
		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}

type UserSign struct {
	Name string `json:"username"`
	Pwd  string `json:"password"`
}

func HandleLogin(a *HApp) {
	a.handler.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		res := &ApiRsp{}
		ret := 0
		cv := ""
		for {
			if r.Body == nil {
				ret = 2
				break
			}
			u := UserSign{}
			e := json.NewDecoder(r.Body).Decode(&u)
			if e != nil {
				res.Msg = e.Error()
				ret = 3
				break
			}
			user := GetUser(a, u.Name)
			if user.Name == "" {
				ret = 1
				break
			}
			if u.Pwd != user.Pwd {
				ret = 5
				break
			}
			cv, e = Encrypt(user.Name+"|"+strconv.FormatInt(time.Now().Unix(), 10), a.key)
			if e != nil {
				res.Msg = e.Error()
				ret = 6
				break
			}
			break
		}

		if ret != 0 {
			res.Code = ret
			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
			return
		}

		cookie := &http.Cookie{
			Name:    "token",
			Value:   cv,
			Expires: time.Now().Add(Token_Expire_Time),
			Path:    "/",
		}
		http.SetCookie(w, cookie)
		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}

func HandleSignUp(a *HApp) {
	a.handler.HandleFunc("/api/signup", func(w http.ResponseWriter, r *http.Request) {
		res := &ApiRsp{}
		ret := 0
		cv := ""
		u := UserSign{}
		for {
			if r.Body == nil {
				ret = 2
				break
			}
			e := json.NewDecoder(r.Body).Decode(&u)
			if e != nil {
				res.Msg = e.Error()
				ret = 3
				break
			}
			if u.Name == "" || u.Pwd == "" {
				ret = 7
				break
			}
			if u.Name != Admin_User_Name && !a.openRegiste { //admin可以注册
				ret = 8
				res.Msg = "registe close"
				break
			}
			user := GetUser(a, u.Name)
			if user.Name != "" {
				ret = 1
				break
			}
			cv, e = Encrypt(u.Name+"|"+strconv.FormatInt(time.Now().Unix(), 10), a.key)
			if e != nil {
				res.Msg = e.Error()
				ret = 6
				break
			}
			break
		}

		if ret != 0 {
			res.Code = ret
			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
			return
		}

		AddUser(a, User{Name: u.Name, Pwd: u.Pwd})

		cookie := &http.Cookie{
			Name:    "token",
			Value:   cv,
			Expires: time.Now().Add(Token_Expire_Time),
			Path:    "/",
		}
		http.SetCookie(w, cookie)
		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}

type UserChpwd struct {
	OldP string `json:"old_password"`
	NewP string `json:"new_password"`
	AgaP string `json:"again_password"`
}

func HandleChpwd(a *HApp) {
	a.handler.HandleFunc("/api/chpwd", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		params := a.ParseHttpParams(r)
		name := params["username"]
		user := GetUser(a, name)
		if name == "" || user.Name == "" {
			w.Write([]byte(LoginHtml))
			return
		}

		res := &ApiRsp{}
		ret := 0
		u := UserChpwd{}
		for {
			if r.Body == nil {
				ret = 1
				break
			}
			json.Unmarshal(body, &u)
			if u.OldP != user.Pwd {
				ret = 2
				break
			}
			if u.NewP == "" {
				ret = 3
				break
			}
			break
		}

		if ret != 0 {
			res.Code = ret
			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
			return
		}

		user.Pwd = u.NewP
		AddUser(a, user)

		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}

type TableResponse struct {
	Code  int64         `json:"code"`
	Count int64         `json:"count"`
	data  []interface{} `json:"data"`
	Msg   string        `json:"msg"`
}

func HandleTable(a *HApp) {
	a.handler.HandleFunc("/api/table", func(w http.ResponseWriter, r *http.Request) {
		params := a.ParseHttpParams(r)
		user := params["username"]
		event_id := params["event_id"]
		url_router := params["url_router"]
		tbnew := a.GetElem(user, url_router, event_id)
		if tbnew == nil {
			io.WriteString(w, `{"code":1,"msg":"no table","count":0,"data":[]}`)
			return
		}
		tb := tbnew.(*HTable)
		oper := params["oper"]
		if oper == "data" {
			page := params["page"]
			limit := params["limit"]
			p, _ := strconv.Atoi(page)
			l, _ := strconv.Atoi(limit)
			sr := params["key[search]"]
			var cnt int
			var data [][]string
			if tb.funcData != nil {
				cnt, data = tb.funcData(user, p, l, sr)
			} else {
				cnt, data = tb.TableGetData(user, p, l, sr)
			}
			retData := ""
			ln := 0
			for _, cols := range data {
				if len(tb.Header) == len(cols) {
					if ln > 0 {
						retData += ","
					}
					cc := "{"
					for i, c := range cols {
						if i > 0 {
							cc += ","
						}
						k := "col" + strconv.Itoa(i)
						kk, _ := json.Marshal(&k)
						v, _ := json.Marshal(&c)
						cc += string(kk) + ":" + string(v)
					}
					cc += "}"
					retData += cc
					ln++
				}
			}
			io.WriteString(w, fmt.Sprintf("{\"code\":0,\"msg\":\"\",\"count\":%d,\"data\":[%s]}", cnt, retData))
		} else if oper == "edit" {
			var res ApiRsp
			if tb.funcEvent != nil {
				res = tb.funcEvent(user, TOEdit, TableAddCols(params))
			} else {
				res = tb.TableEvent(user, TOEdit, TableAddCols(params))
			}

			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
		} else if oper == "del" {
			var res ApiRsp
			if tb.funcEvent != nil {
				res = tb.funcEvent(user, TODel, TableAddCols(params))
			} else {
				res = tb.TableEvent(user, TODel, TableAddCols(params))
			}

			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
		} else if oper == "add" && tb.funcEvent != nil {
			var res ApiRsp
			if tb.funcEvent != nil {
				res = tb.funcEvent(user, TOAdd, TableAddCols(params))
			} else {
				res = tb.TableEvent(user, TOAdd, TableAddCols(params))
			}

			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
		} else if oper == "url" {
			if tb.funcUrl != nil {
				href := params["href"]
				io.WriteString(w, tb.funcUrl(user, href))
			} else {
				io.WriteString(w, Html404)
			}
		} else if strings.HasPrefix(oper, "user_") {
			res := &HResponse{Error: "error param"}
			ev, e := strconv.Atoi(oper[5:])
			if e == nil && tb.funcUsrEvent != nil {
				res = tb.funcUsrEvent(user, ev, TableAddCols(params))
			}

			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
		} else {
			io.WriteString(w, `{"code":1,"msg":"error params","count":0,"data":[]}`)
		}
	})
}

func HandleMergelay(a *HApp) {
	a.handler.HandleFunc("/api/mergely", func(w http.ResponseWriter, r *http.Request) {
		params := a.ParseHttpParams(r)
		user := params["username"]
		event_id := params["event_id"]
		url_router := params["url_router"]
		file := params["file"]

		res := &ApiRsp{}
		defer func() {
			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
		}()

		h := a.GetElem(user, url_router, event_id)
		if h == nil {
			res.Code = 1
			res.Msg = "no mergely item id=" + event_id
			return
		}
		mer := h.(*HMergely)

		if mer.F == nil {
			res.Code = 1
			res.Msg = "you should overload HMergely.OnGetFile"
		} else {
			res.Msg = mer.F(user, file)
		}
	})
}

func HandleUpload(a *HApp) {
	a.handler.HandleFunc("/api/upload", func(w http.ResponseWriter, r *http.Request) {
		res := &HResponse{}
		defer func() {
			ret_json, _ := json.Marshal(res)
			io.WriteString(w, string(ret_json))
		}()

		params := a.ParseHttpParams(r)
		user := params["username"]
		event_id := params["event_id"]
		url_router := params["url_router"]
		res.RedirectUrl = url_router
		up := a.GetElem(user, url_router, event_id)
		if up == nil {
			res.Error = "no upload item"
			return
		}
		u1 := up.(*HUpload)
		var fileName string
		var fileData []byte

		reader, err := r.MultipartReader()
		if err != nil {
			res.Error = err.Error()
			return
		}

		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				res.Error = err.Error()
				return
			}
			if part.FileName() == "" { // this is FormData
				//data, _ := ioutil.ReadAll(part)
				//fmt.Printf("FormData=[%s]\n", string(data))
			} else { // This is FileData
				fileName = part.FileName()
				data, _ := ioutil.ReadAll(part)
				fileData = append(fileData, data...)
			}
		}

		if u1.OnUpload != nil {
			u1.OnUpload(user, params, fileName, fileData)
		}
	})
}
