// event.go
package ssui

import (
	"encoding/json"
	"io"
	"net/http"
)

/**************************************************************************
			event_id 事件发起者ID
			url_router 事件页面路由
			token 认证
			其他key val分别表示页面控件ID和Val
**************************************************************************/
func HandleButtonClick(a *HApp) {
	a.Handler.HandleFunc("/button_click", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		params := make(map[string]string, 0)
		for k, v := range r.Form {
			params[k] = v[0]
		}
		event_id, ok := params["event_id"]
		url_router, ok1 := params["url_router"]
		token := params["token"]
		res := &HResponse{Error: "error param", Token: token}
		for {
			if !ok || !ok1 {
				break
			}
			e := a.GetElem(token, url_router, event_id)
			if e == nil {
				break
			}
			b := e.(*HButton)
			if b.Event == nil {
				break
			}
			res = b.Event(params)
			break
		}

		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}

/**************************************************************************
			event_id 事件发起者ID
			url_router 事件页面路由
			token 认证
			rowid 行ID
**************************************************************************/
func HandleTableDel(a *HApp) {
	a.Handler.HandleFunc("/table_del", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		params := make(map[string]string, 0)
		for k, v := range r.Form {
			params[k] = v[0]
		}
		event_id, ok := params["event_id"]
		url_router, ok1 := params["url_router"]
		token := params["token"]
		res := &HResponse{Error: "error param", Token: token}
		for {
			if !ok || !ok1 {
				break
			}
			e := a.GetElem(token, url_router, event_id)
			if e == nil {
				break
			}
			t := e.(*HTable)
			if t.EventDel == nil {
				break
			}
			res = t.EventDel(t, TableDelRowId(params))
			break
		}

		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}

/**************************************************************************
			event_id 事件发起者ID
			url_router 事件页面路由
			token 认证
			0 1 2 3... N列的值
**************************************************************************/
func HandleTableAdd(a *HApp) {
	a.Handler.HandleFunc("/table_add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		params := make(map[string]string, 0)
		for k, v := range r.Form {
			params[k] = v[0]
		}
		event_id, ok := params["event_id"]
		url_router, ok1 := params["url_router"]
		token := params["token"]
		res := &HResponse{Error: "error param", Token: token}
		for {
			if !ok || !ok1 {
				break
			}
			e := a.GetElem(token, url_router, event_id)
			if e == nil {
				break
			}
			t := e.(*HTable)
			if t.EventAdd == nil {
				break
			}
			res = t.EventAdd(t, TableAddCols(params))
			break
		}

		ret_json, _ := json.Marshal(res)
		io.WriteString(w, string(ret_json))
	})
}
