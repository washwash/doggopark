package api

import (
  "net/http"
)


func (this RequestHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	handler := http.DefaultServeMux
	handler.ServeHTTP(rw, req)
}


func Run() {
	handler := RequestHandler{}
	http.ListenAndServe(":8080", handler)
}



