// response.go
package ssui

type HResponse struct {
	Error        string //优先弹出错误
	EchoText     string //次优显示返回内容
	RedirectUrl  string //最后执行跳转页面
	ShowInDialog bool   //跳转页面是否展示在弹出框里
	Token        string //token
}

func ResponseError(err, token string) *HResponse {
	return &HResponse{Error: err, Token: token}
}

func ResponseMsg(err, token string) *HResponse {
	return &HResponse{Error: err, Token: token}
}

func ResponseText(txt, token string) *HResponse {
	return &HResponse{EchoText: txt, Token: token}
}

func ResponseURL(url, token string) *HResponse {
	return &HResponse{RedirectUrl: url, Token: token, ShowInDialog: false}
}

func ResponseDialog(url, token string) *HResponse {
	return &HResponse{RedirectUrl: url, Token: token, ShowInDialog: true}
}
