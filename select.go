// select.go
package ssui

type HSelect struct {
	*ElemBase
	SelIndex int
	Option   []string
}

var HtmlSelect = `<div class="layui-form-item {{if .Hide}}layui-hide{{end}}">
<select id="{{.Id}}" lay-filter="{{.Id}}">
{{range $i,$v:=.Option}}<option {{if eq $i $.SelIndex}}selected{{end}} {{if $.Disable}}disabled=""{{end}} value="{{$i}}">{{$v}}</option>{{end}}
</select>
<script>
	layui.use(['form'], function () {
	     var $ = layui.jquery,
	         form = layui.form;
	    
		form.on('select({{.Id}})', function (data) {
	       //console.log(this.innerText);
	       //console.log(data.value);
	   	});
	});
</script>
</div>`

func NewSelect(id string, selindex int, options []string) *HSelect {
	no := make([]string, len(options), len(options))
	copy(no, options)
	s := &HSelect{newElem(id, "select", HtmlSelect), selindex, no}
	s.self = s
	return s
}
func (s *HSelect) Clone() HtmlElem {
	ns := NewSelect(s.Id, s.SelIndex, s.Option)
	ns.ElemBase.clone(s.ElemBase)
	return ns
}
