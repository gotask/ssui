// echart
package ssui

type XData struct {
	Title string
	X     []int
}

type HEchart struct {
	*ElemBase
	Title string
	Fill  bool
	X     []string
	D     []XData
}

func (e *HEchart) SetX(x []string) {
	nx := make([]string, len(x), len(x))
	copy(nx, x)
	e.X = nx
}
func (e *HEchart) AddData(title string, xdata []int) {
	e.D = append(e.D, XData{title, xdata})
}

func (e *HEchart) Reset() {
	e.X = nil
	e.D = nil
}

func (e *HEchart) SetFill(f bool) {
	e.Fill = f
}

var HtmlEchart = `<div id="{{.Id}}" {{if .Hide}}class="layui-hide"{{end}} style="min-height:300px;padding: 10px">
<script>
	layui.use(['echarts'], function () {
	     var $ = layui.jquery,
	         echarts = layui.echarts;
	    
        var echartsRecords_{{RawString .Id}} = echarts.init(document.getElementById('{{.Id}}'), 'walden');

        var optionRecords_{{RawString .Id}} = {
            title: {
                text: '{{.Title}}'
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985'
                    }
                }
            },
            legend: {
               data: [{{range $i,$v:=.D}} {{if gt $i 0}},{{end}}'{{$v.Title}}'{{end}}]
            },
            toolbox: {
                feature: {
                    saveAsImage: {}
                }
            },
            grid: {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: [{{range $i,$v:=.X}} {{if gt $i 0}},{{end}}'{{$v}}'{{end}}]
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
				{{range $i,$v:=.D}} {{if gt $i 0}},{{end}}{name: '{{$v.Title}}',type: 'line',{{if $.Fill}}areaStyle: {},{{end}}
				data:[{{range $i,$v:=.X}}{{if gt $i 0}},{{end}}{{RawInt $v}}{{end}}]}{{end}}
            ]
        };
        echartsRecords_{{RawString .Id}}.setOption(optionRecords_{{RawString .Id}});
		// echarts 窗口缩放自适应
		$(window).resize(function(){
		    echartsRecords_{{RawString .Id}}.resize();
		});
	});
</script>
</div>
`

func NewEchart(id, title string) *HEchart {
	e := &HEchart{newElem(id, "echart", HtmlEchart), title, false, nil, nil}
	e.self = e
	return e
}
func (s *HEchart) Clone() HtmlElem {
	e := NewEchart(s.Id, s.Title)
	nx := make([]string, len(s.X), len(s.X))
	copy(nx, s.X)
	e.X = nx
	nd := make([]XData, len(s.D), len(s.D))
	copy(nd, s.D)
	e.D = nd
	e.Fill = s.Fill
	e.ElemBase.clone(s.ElemBase)
	return e
}
