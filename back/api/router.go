package api

import (
  "net/http"
)


type Route struct {
	Url string
	Controller http.HandlerFunc
}


func AddRoute(url string, controller HandlerFunc) {
	var c = func(w http.ResponseWriter, r *http.Request){
		req := IRequest(r)
		controller(w, &req)
	}
	http.HandleFunc(url, c)
}


type HandlerFunc = func (IResponseWriter, *IRequest)
type IResponseWriter interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(int)
}
type IRequest interface {}
