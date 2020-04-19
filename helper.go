// helper.go
package ssui

import (
	"strconv"
)

func Value(id string, param map[string]string) string {
	if v, ok := param[id]; ok {
		return v
	}
	return ""
}

func Checked(id string, param map[string]string) bool {
	v := Value(id, param)
	if v == "1" {
		return true
	}
	return false
}

func RadioIndex(id string, param map[string]string) int {
	v := Value(id, param)
	i, _ := strconv.Atoi(v)
	return i
}

func Router(param map[string]string) string {
	return Value("url_router", param)
}

func Sender(param map[string]string) string {
	return Value("event_id", param)
}

func TableDelRowId(param map[string]string) string {
	return Value("rowid", param)
}

func TableAddCols(param map[string]string) []string {
	cols := make([]string, 0)
	for i := 0; i < 1000; i++ {
		is := strconv.Itoa(i)
		if n, ok := param[is]; ok {
			cols = append(cols, n)
		} else {
			return cols
		}
	}
	return cols
}
