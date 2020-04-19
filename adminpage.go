// adminpage
package ssui

func UserSetting(a *HApp) {
	if !a.Admin {
		return
	}
	frameUserSetting := NewFrame("/usersetting", "UserSetting", "layui-icon layui-icon-username", func(user string) {
		u := GetUser(a, user)
		a.GetElem(user, "/usersetting", "phone").(*HLineEdit).Text = u.Phone
		a.GetElem(user, "/usersetting", "remark").(*HTextArea).Text = u.Remark
	})
	frameUserSetting.AddElem(NewRow().AddElem(NewLabel("手机号")).AddElem(NewLineEdit("phone", "", "", false)))
	frameUserSetting.AddElem(NewLabel("备注")).AddElem(NewTextArea("remark", "", ""))
	frameUserSetting.AddElem(NewButton("save", "保存", func(params map[string]string) *HResponse {
		user := params["username"]
		u := GetUser(a, user)
		if u.Name != "" {
			u.Phone = params["phone"]
			u.Remark = params["remark"]
			AddUser(a, u)
			return &HResponse{"保存成功", "", "", false}
		}
		return &HResponse{"保存失败", "", "", false}
	}))
	a.AddFrame(frameUserSetting)
}

func AuthEdit(a *HApp) {
	if !a.Admin {
		return
	}
	authFrame := NewFrame("/authedit", "Auth", "layui-icon layui-icon-set-fill", nil)
	authFrame.AddElem(NewToolTable("users", false, []string{"Name", "Auth", "Phone", "Remark"}, func(user string, page, limit int, searchtxt string) (total int, data [][]string) {
		all := AllUser()
		data = make([][]string, 0, 0)
		begin := (page - 1) * limit
		end := begin + limit
		for i := begin; i >= 0 && i < len(all) && i < end; i++ {
			u := GetUser(a, all[i])
			data = append(data, []string{u.Name, u.Auth, u.Phone, u.Remark})
		}
		return len(all), data
	}, func(user string, t TableOperType, cols []string) ApiRsp {
		if len(cols) != 4 {
			return ApiRsp{1, "error param"}
		}
		if t == TOEdit {
			u := GetUser(a, cols[0])
			if u.Name != "" {
				u.Auth = cols[1]
				u.Phone = cols[2]
				u.Remark = cols[3]
				AddUser(a, u)
			}
		} else if t == TOAdd {
			u := GetUser(a, cols[0])
			if u.Name != "" {
				return ApiRsp{1, "用户已存在"}
			}
			u.Name = cols[0]
			u.Pwd = "12345"
			u.Auth = cols[1]
			u.Phone = cols[2]
			u.Remark = cols[3]
			AddUser(a, u)
		} else if t == TODel {
			return ApiRsp{1, "不可删除"}
		}
		return ApiRsp{0, ""}
	}))
	gp := NewPageGroup("AuthEdit", "layui-icon layui-icon-group")
	gp.AddFrame(authFrame)
	a.AddPageGroup(gp)
}
