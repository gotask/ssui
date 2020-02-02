# ssui
a simper html ui library for freshman

#example
```
app := NewApp(":8080")
f := NewFrame("/", "HL", "Hello", "World")
f.AddElem(NewRow().AddElem(NewLabel("Label1")).AddElem(NewLineEdit("LineEdit1", "input", "", false)))
f.AddElem(NewButton("Button1", "Click", func(param map[string]string) *HResponse {
	if Checked("CheckBox1", param) && RadioIndex("Radio1", param) == 1 &&
		Value("Select2", param) == "f" {
		GetAppElem("Button1", app, param).(*HButton).Text = Value("LineEdit1", param)
	}
	return &HResponse{"", "", "/", true}
}))
f.AddElem(NewTextArea("TextArea1", "input...", "")).AddElem(NewSelect("Select1", 0, []string{"a", "b", "c"}))
f.AddElem(NewSelect("Select2", 1, []string{"d", "f", "e"}))
f.AddElem(NewRadio("Radio1", 0, []string{"a", "b", "c"}))
f.AddElem(NewTimePicker("TimePicker1", "yyyy-MM-dd HH:mm:ss", 0))
f.AddElem(NewRow().AddElem(NewCheckBox("CheckBox1", "a", true)).AddElem(NewCheckBox("CheckBox2", "b", true)).AddElem(NewCheckBox("CheckBox3", "c", false)))

table := &HTable{"Table1", []string{"a", "b", "c"}, []HTableRow{
	HTableRow{"3", []HTableElem{HTableElem{"Hello"},
		HTableElem{"world"},
		HTableElem{"golang"}}},
	HTableRow{"4", []HTableElem{HTableElem{"Hello"},
		HTableElem{"world"},
		HTableElem{"golang"}}}}, true, true, func(t *HTable, rowid string) *HResponse {
	for i, r := range t.Rows {
		if r.Id == rowid {
			t.Rows = append(t.Rows[0:i], t.Rows[i+1:]...)
			break
		}
	}
	return &HResponse{"", "", "/", false}
}, func(t *HTable, cols []string) *HResponse {
	if len(cols) == 0 || cols[0] == "" {
		return &HResponse{"table id error", "", "", false}
	}
	newrow := HTableRow{}
	newrow.Id = cols[0]
	for _, c := range cols {
		newrow.Elems = append(newrow.Elems, HTableElem{c})
	}
	t.Rows = append(t.Rows, newrow)
	return &HResponse{"", "", "/", false}
}}
f.AddElem(table)

app.AddFrame(f).Run()
```
