package ssui

type HBlank struct {
	*ElemBase
}

var HtmlBlank = `<div class="layui-form-item">
<label class="layui-form-label" style="opacity: 0;">-</label>
</div>`

func NewBlank() *HBlank {
	return &HBlank{newElem("", "blank", HtmlBlank)}
}

func (l *HBlank) Clone() HtmlElem {
	nl := NewBlank()
	nl.ElemBase.clone(l.ElemBase)
	return nl
}
