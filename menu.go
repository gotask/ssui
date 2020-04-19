// menuinfo.go
package ssui

type Menu struct {
	HomeInfo MenuHomeInfo `json:"homeInfo"`
	LogoInfo MenuLogoInfo `json:"logoInfo"`
	MenuInfo []MenuChild  `json:"menuInfo"`
}

type MenuHomeInfo struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

type MenuLogoInfo struct {
	Href  string `json:"href"`
	Image string `json:"image"`
	Title string `json:"title"`
}

type MenuChild struct {
	Child  []MenuChild `json:"child"`
	Href   string      `json:"href"`
	Icon   string      `json:"icon"`
	Target string      `json:"target"`
	Title  string      `json:"title"`
}
