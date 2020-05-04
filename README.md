# ssui
a simper html ui library for freshman

# example
```
// ui.go
package main

import (
	"strings"

	"github.com/gotask/gost/stconfig"
	"github.com/gotask/gost/stlog"
	"github.com/gotask/gost/stutil"
	"github.com/gotask/ssui"
)

func Display() error {
	LOG := stlog.NewFileLogger("sync.log")
	defer LOG.Close()
	c, e := stconfig.LoadINI("sync.ini")
	if e != nil {
		stutil.FileCreateAndWrite("sync.ini", `[system]
loopsec = 1
address = 0.0.0.0:3030
uiaddr = 0.0.0.0:2020
token = 

[httpdownload1]
dir = .
address = 0.0.0.0:2021

[sync1]
dir = D:\myproject=>/home/xxx/myproject@192.168.1.111:3030
include = .*\.(cpp|h|go)
exclude = .*
completex = build
`)
		c, e = stconfig.LoadINI("sync.ini")
		if e != nil {
			return e
		}
	}

	uiaddr := c.StringSection("system", "uiaddr", "")
	if uiaddr == "" {
		return nil
	}
	token := c.StringSection("system", "token", "")

	var app *ssui.HApp
	router := "/"
	if token != "" {
		app = ssui.NewAdminApp(uiaddr, "gosync", token)
		router = "/config"
	} else {
		app = ssui.NewApp(uiaddr)
	}
	f := ssui.NewFrame(router, "gosync", "layui-icon layui-icon-set", func(user string) {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return
		}
		app.GetElem(user, router, "loopsec").(*ssui.HLineEdit).Text = c.StringSection("system", "loopsec", "1")
		app.GetElem(user, router, "address").(*ssui.HLineEdit).Text = c.StringSection("system", "address", ":3030")
	})
	f.AddElem(ssui.NewLegend("HTTP"))
	f.AddElem(ssui.NewLabel("HTTP下载配置表"))
	f.AddElem(ssui.NewToolTable("httpdownload", false, []string{"ID(格式不能变)", "下载目录", "http地址"}, func(user string, page, limit int, searchtxt string) (total int, data [][]string) {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return 0, nil
		}
		data = make([][]string, 0, 0)
		for i := 1; i <= 100; i++ {
			sec := "httpdownload" + stutil.IntToString(int64(i))
			tempD := c.StringSection(sec, "dir", "")
			if tempD == "" {
				continue
			}
			tempA := c.StringSection(sec, "address", "")
			if tempA == "" {
				continue
			}
			data = append(data, []string{sec, tempD, tempA})
		}
		return len(data), data
	}, func(user string, t ssui.TableOperType, cols []string) ssui.ApiRsp {
		if len(cols) != 3 {
			return ssui.ApiRsp{1, "error param"}
		}
		if t == ssui.TOEdit || t == ssui.TOAdd {
			c, e := stconfig.LoadINI("sync.ini")
			if e != nil {
				return ssui.ApiRsp{1, e.Error()}
			}
			if !strings.HasPrefix(cols[0], "httpdownload") || cols[1] == "" || cols[2] == "" {
				return ssui.ApiRsp{1, "error param"}
			}
			c.SectionSet(cols[0], "dir", cols[1], "")
			c.SectionSet(cols[0], "address", cols[2], "")
			c.Save()
		} else if t == ssui.TODel {
			c, e := stconfig.LoadINI("sync.ini")
			if e != nil {
				return ssui.ApiRsp{1, e.Error()}
			}
			c.DelSection(cols[0])
			c.Save()
		}
		return ssui.ApiRsp{0, ""}
	}))
	f.AddElem(ssui.NewLegend("SYNC"))
	f.AddElem(ssui.NewRow().AddElem(ssui.NewLabel("sync同步配置表")).AddElem(ssui.NewButton("resync", "重新同步", func(param map[string]string) *ssui.HResponse {
		return ssui.ResponseError("开始同步，请观察日志...")
	})).AddElem(ssui.NewLabel("循环间隔")).AddElem(ssui.NewLineEdit("loopsec", "s", c.StringSection("system", "loopsec", "1"), false)).AddElem(
		ssui.NewLabel("监听地址")).AddElem(ssui.NewLineEdit("address", ":3030", c.StringSection("system", "address", ":3030"), false)).AddElem(ssui.NewButton("change", "提交", func(param map[string]string) *ssui.HResponse {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return ssui.ResponseError(e.Error())
		}
		loopsec := ssui.Value("loopsec", param)
		address := ssui.Value("address", param)
		if loopsec == "" || address == "" {
			return ssui.ResponseError("error param")
		}
		c.SectionSet("system", "loopsec", loopsec, "")
		c.SectionSet("system", "address", address, "")
		c.Save()
		return ssui.ResponseError("保存成功，重启生效")
	})))
	f.AddElem(ssui.NewToolTable("sync", false, []string{"ID(格式不能变)", "同步地址", "Include", "Exclude", "CompleteExclude"}, func(user string, page, limit int, searchtxt string) (total int, data [][]string) {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return 0, nil
		}
		data = make([][]string, 0, 0)
		for i := 1; i <= 100; i++ {
			sec := "sync" + stutil.IntToString(int64(i))
			di := c.StringSection(sec, "dir", "")
			if di == "" {
				continue
			}
			in := c.StringSection(sec, "include", "")
			ex := c.StringSection(sec, "exclude", "")
			cex := c.StringSection(sec, "completex", "")
			data = append(data, []string{sec, di, in, ex, cex})
		}
		return len(data), data
	}, func(user string, t ssui.TableOperType, cols []string) ssui.ApiRsp {
		if len(cols) != 5 {
			return ssui.ApiRsp{1, "error param"}
		}
		if t == ssui.TOEdit || t == ssui.TOAdd {
			c, e := stconfig.LoadINI("sync.ini")
			if e != nil {
				return ssui.ApiRsp{1, e.Error()}
			}
			if !strings.HasPrefix(cols[0], "sync") || cols[1] == "" {
				return ssui.ApiRsp{1, "error param"}
			}
			c.SectionSet(cols[0], "dir", cols[1], "")
			c.SectionSet(cols[0], "include", cols[2], "")
			c.SectionSet(cols[0], "exclude", cols[3], "")
			c.SectionSet(cols[0], "completex", cols[4], "")
			c.Save()
		} else if t == ssui.TODel {
			c, e := stconfig.LoadINI("sync.ini")
			if e != nil {
				return ssui.ApiRsp{1, e.Error()}
			}
			c.DelSection(cols[0])
			c.Save()
		}
		return ssui.ApiRsp{0, ""}
	}))
	f.AddElem(ssui.NewLegend("LOG"))
	f.AddElem(ssui.NewButton("log", "日志", func(param map[string]string) *ssui.HResponse {
		buf, e := stutil.FileReadAll("sync.log")
		if e != nil {
			return ssui.ResponseError(e.Error())
		}
		ret := ""
		if len(buf) < 1024 {
			ret = string(buf)
		} else {
			ret = string(buf[len(buf)-1024:])
		}
		app.GetElem(param["username"], router, "code").(*ssui.HText).Text = ret
		return ssui.ResponseURL(router)
	}))
	f.AddElem(ssui.NewText("code", ""))

	if token != "" {
		gp := ssui.NewPageGroup("Config", "layui-icon layui-icon-set-fill")
		gp.AddFrame(f)
		app.AddPageGroup(gp)
	} else {
		app.AddFrame(f)
	}
	err := app.Run()
	if err != nil {
		LOG.Error("ui start error: %s", err.Error())
	} else {
		LOG.Info("ui start success: %s", uiaddr)
	}
	return err
}
```
![效果](https://github.com/gotask/images/blob/master/gosync.jpg)