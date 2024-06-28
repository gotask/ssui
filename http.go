package ssui

import (
	"github.com/gotask/gost/stnet"
	"net/http"
)

type HttpServer struct {
}

func (hp *HttpServer) Init() bool {
	return true
}
func (hp *HttpServer) Loop() {

}
func (hp *HttpServer) HandleError(current *stnet.CurrentContent, e error) {
	if current != nil && current.Sess != nil {
		current.Sess.Close()
	}
}
func (hp *HttpServer) HashProcessor(current *stnet.CurrentContent, req *http.Request) (processorID int) {
	return -1
}
