// table.go
package ssui

import (
	"bytes"
	"html/template"
)

type OnTableEventDel func(t *HTable, rowid string) *HResponse
type OnTableEventAdd func(t *HTable, cols []string) *HResponse

type HTableElem struct {
	Text string
}
type HTableRow struct {
	Id    string
	Elems []HTableElem
}
type HTable struct {
	Id       string
	Header   []string
	Rows     []HTableRow
	Del      bool
	Add      bool
	EventDel OnTableEventDel
	EventAdd OnTableEventAdd
}

var HtmlTable = `
<table class="layui-table layui-text" id="{{.Id}}">
    <thead>
      <tr>
	{{range .Header}} <th>{{.}}</th> {{end}}
	{{if .Del}} <th>Tool</th> {{end}}
      </tr> 
    </thead>
 <tbody>{{$del:=.Del}}{{$add:=.Add}}{{$TId:=.Id}}
	{{range $i,$v:=.Rows}}<tr>
		{{range $v.Elems}}<td>{{.Text}}</td>{{end}} 
		{{if $del}} <td><a class="layui-btn" onclick="tableDel('{{$TId}}', '{{.Id}}')">Del</a>{{if $add}}<a class="layui-btn" onclick="tableCpy('{{$TId}}', {{$i}})">Cpy</a>{{end}} </td>{{end}}
	</tr>{{end}}
 </tbody>{{if .Add}} 
    <thead>
	<tr>
	{{range $k,$v:=.Header}} <th><input class="layui-input" id="{{$TId}}_{{$k}}"></input></th> {{end}}
	<th><button type="button" class="layui-btn" id="TableAdd_{{$TId}}">Add</button></th>
    </tr> 
    </thead>
	<script>
	layui.use(['form'], function () {
	        var $ = layui.jquery,
	            form = layui.form;
	
	 	$("#TableAdd_{{$TId}}").click(function() {
			var $ = layui.jquery;
			var url = "/table_add?event_id={{$TId}}&url_router="+window.location.pathname+"&token="+getToken(){{range $k,$v:=.Header}}+"&{{$k}}="+$("#{{$TId}}_{{$k}}").val(){{end}};
			$.get(url,function(ret){
				handleRsp(ret);
		    });
		});
	});
	</script>{{end}}
</table>

`

func NewTable(id string, header []string, del, add bool, eventdel OnTableEventDel, evenadd OnTableEventAdd) *HTable {
	h := make([]string, len(header), len(header))
	copy(h, header)
	return &HTable{id, h, nil, del, add, eventdel, evenadd}
}

func (table *HTable) Type() string {
	return "table"
}
func (table *HTable) ID() string {
	return table.Id
}
func (table *HTable) Clone() HtmlElem {
	nt := NewTable(table.Id, table.Header, table.Del, table.Add, table.EventDel, table.EventAdd)
	for _, r := range table.Rows {
		rn := HTableRow{}
		rn.Id = r.Id
		for _, e := range r.Elems {
			rn.Elems = append(rn.Elems, e)
		}
		nt.Rows = append(nt.Rows, rn)
	}
	return nt
}
func (table *HTable) Render(token string) string {
	te := template.New("table")
	t, e := te.Parse(HtmlTable)
	if e != nil {
		return e.Error()
	}
	buf := bytes.NewBufferString("")
	e = t.Execute(buf, table)
	if e != nil {
		return e.Error()
	}
	return buf.String()
}

func (table *HTable) AddRow(id string, cols []string) {
	row := HTableRow{id, make([]HTableElem, 0)}
	for _, c := range cols {
		row.Elems = append(row.Elems, HTableElem{c})
	}
	table.Rows = append(table.Rows, row)
}

func (table *HTable) Reset() {
	table.Rows = nil
}
