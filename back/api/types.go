package api

import (
  "net/http"
)


type (
	Context struct {
		Request http.Request
	}

	Controller func (c Context) (interface{}, int)

	RequestHandler struct {}
)
