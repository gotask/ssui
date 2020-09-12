package ssui

import "strconv"

type HProgress struct {
	*ElemBase
	Percent float32
}

var HtmlProgress = `<div class="layui-progress {{if .Hide}}layui-hide{{end}}" lay-filter="{{.Id}}">
  <input type="text" class="layui-hide" id="{{.Id}}" value="{{.Percent}}">
  <div class="layui-progress-bar layui-bg-red" lay-percent="{{.Percent}}%"></div>
<script>
	layui.use(['form','element'], function () {
	     var $ = layui.jquery,
			 form = layui.form,
	         element = layui.element;
	});
</script>
</div>`

func NewProgress(id string, percent float32) *HProgress {
	p := &HProgress{newElem(id, "progress", HtmlProgress), percent}
	p.self = p
	return p
}

func (p *HProgress) Clone() HtmlElem {
	np := NewProgress(p.Id, p.Percent)
	np.ElemBase.clone(p.ElemBase)
	return np
}

func (p *HProgress) SetValue(v string) {
	f, e := strconv.ParseFloat(v, 32)
	if e != nil {
		return
	}
	p.Percent = float32(f)
}
