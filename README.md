# ssui
a simper html ui library for freshman

# example
```
package main

import (
	"fmt"
	"strconv"

	"github.com/gotask/ssui"
)

func main() {
	fmt.Println("Hello World!")
	app := ssui.NewApp(":1010")
	login := ssui.NewFrame("/login", "gosync", "gosync", "gosync", func(token string) {
	}, func(token string) (success bool, tokenFailedUrl string) {
		return true, ""
	})
	login.AddElem(ssui.NewLabel("Token")).AddElem(ssui.NewLineEdit("loginToken", "token", "", false)).AddElem(ssui.NewButton("loginReq", "验证", func(param map[string]string) *ssui.HResponse {
		token := ssui.Value("loginToken", param)
		if "12345" != token {
			return ssui.ResponseError("验证失败", "")
		}
		return ssui.ResponseURL("/", token)
	}))
	f := ssui.NewFrame("/", "gosync", "gosync", "gosync", func(token string) {
		app.GetElem(token, "/", "loopsec").(*ssui.HLineEdit).Text = "1"
		app.GetElem(token, "/", "address").(*ssui.HLineEdit).Text = ":3030"
		httptable := app.GetElem(token, "/", "httpdownload").(*ssui.HTable)
		httptable.Reset()
		for i := 1; i <= 3; i++ {
			httptable.AddRow(strconv.Itoa(i), []string{strconv.Itoa(i), "D", "A"})
		}

		synctable := app.GetElem(token, "/", "sync").(*ssui.HTable)
		synctable.Reset()
		for i := 1; i <= 3; i++ {
			synctable.AddRow(strconv.Itoa(i), []string{strconv.Itoa(i), "D", "1", "2", "3"})
		}
	}, func(token string) (success bool, tokenFailedUrl string) {
		if "12345" == token {
			return true, ""
		}
		return false, "/login"
	})
	f.AddElem(ssui.NewLegend("HTTP"))
	f.AddElem(ssui.NewLabel("HTTP下载配置表"))
	f.AddElem(ssui.NewTable("httpdownload", []string{"ID(格式不能变)", "下载目录", "http地址"}, true, true, func(t *ssui.HTable, rowid string) *ssui.HResponse {
		return ssui.ResponseURL("/", "")
	}, func(t *ssui.HTable, cols []string) *ssui.HResponse {
		return ssui.ResponseURL("/", "")
	}))
	f.AddElem(ssui.NewLegend("SYNC"))
	f.AddElem(ssui.NewRow().AddElem(ssui.NewLabel("sync同步配置表")).AddElem(ssui.NewButton("resync", "重新同步", func(param map[string]string) *ssui.HResponse {
		return ssui.ResponseMsg("开始同步，请观察日志...", "")
	})).AddElem(ssui.NewLabel("循环间隔")).AddElem(ssui.NewLineEdit("loopsec", "s", "", false)).AddElem(
		ssui.NewLabel("监听地址")).AddElem(ssui.NewLineEdit("address", ":3030", "", false)).AddElem(ssui.NewButton("change", "提交", func(param map[string]string) *ssui.HResponse {
		return ssui.ResponseMsg("保存成功，重启生效", "")
	})))
	f.AddElem(ssui.NewTable("sync", []string{"ID(格式不能变)", "同步地址", "Include", "Exclude", "CompleteExclude"}, true, true, func(t *ssui.HTable, rowid string) *ssui.HResponse {
		return ssui.ResponseURL("/", "")
	}, func(t *ssui.HTable, cols []string) *ssui.HResponse {
		return ssui.ResponseURL("/", "")
	}))
	f.AddElem(ssui.NewLegend("LOG"))
	f.AddElem(ssui.NewButton("log", "日志", func(param map[string]string) *ssui.HResponse {
		return ssui.ResponseText("test", "")
	}))
	app.AddFrame(f).AddFrame(login).Run()
}
```
![效果](https://github.com/gotask/images/blob/master/ssui.png)