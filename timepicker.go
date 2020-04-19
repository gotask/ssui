// timepicker.go
package ssui

type HTimePicker struct {
	*ElemBase
	//y代表年M代表月,以此类推,例如: yyyy-MM-dd HH:mm:ss yyyy年M月 yyyy年的M月某天晚上，大概H点 dd/MM/yyyy ...
	Format string
	Value  int64 //1970年以来的ms数
}

var HtmlTimePicker = `<div class="layui-form-item">
<input type="text" class="layui-input" id="{{.Id}}">
<script>
	layui.use(['form','laydate'], function () {
	     var $ = layui.jquery,
	         form = layui.form,
			 laydate = layui.laydate;

	    //日期
        laydate.render({
            elem: '#{{.Id}}'
			,type: 'datetime'
            ,format: '{{.Format}}'
            ,value: new Date({{if gt .Value 0}}{{.Value}}{{end}})
            ,isInitValue: true
            ,done: function(value, date, endDate){
			    //console.log(value); //得到日期生成的值，如：2017-08-18
			}
        });
	});
</script>
</div>`

//y代表年M代表月,以此类推,例如: yyyy-MM-dd HH:mm:ss yyyy年M月 yyyy年的M月某天晚上，大概H点 dd/MM/yyyy ... 1970年以来的ms数
func NewTimePicker(id, format string, val int64) *HTimePicker {
	p := &HTimePicker{newElem(id, "timepicker", HtmlTimePicker), format, val}
	p.self = p
	return p
}
func (p *HTimePicker) Clone() HtmlElem {
	np := NewTimePicker(p.Id, p.Format, p.Value)
	np.ElemBase.clone(p.ElemBase)
	return np
}
