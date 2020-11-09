// response.go
package ssui

type ApiRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type HResponse struct {
	Error        string //弹出错误
	RedirectUrl  string //执行跳转页面
	ShowInDialog string //跳转页面是否展示在弹出框里 弹出框标题
	ShowHtml     string //直接弹窗显示返回的string
	SelfClose    bool   //是否关闭当前页面
}

func ResponseError(err string) *HResponse {
	return &HResponse{Error: err}
}

func ResponseURL(url string) *HResponse {
	return &HResponse{RedirectUrl: url, ShowInDialog: "", SelfClose: true}
}

func ResponseDialog(url, title string) *HResponse {
	return &HResponse{RedirectUrl: url, ShowInDialog: title}
}

func ResponseHtml(html, title string) *HResponse {
	return &HResponse{ShowHtml: `<pre style="word-wrap: break-word; white-space: pre-wrap;">` + html + "</pre>", ShowInDialog: title}
}
