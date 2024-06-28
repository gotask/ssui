// pagegroup.go
package ssui

type PageGroup struct {
	Title  string
	Icon   string
	Frames []*Frame
	Groups []*PageGroup
}

// icon https://layui.dev/2.7/docs/element/icon.html
func NewPageGroup(title, icon string) *PageGroup {
	return &PageGroup{title, icon, make([]*Frame, 0, 0), make([]*PageGroup, 0, 0)}
}

func (p *PageGroup) AddFrame(f *Frame) *PageGroup {
	p.Frames = append(p.Frames, f)
	return p
}

func (p *PageGroup) AddGroup(np *PageGroup) *PageGroup {
	p.Groups = append(p.Groups, np)
	return p
}
