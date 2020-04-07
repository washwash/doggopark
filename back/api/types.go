package api

import (
  "net/http"
)


type Context struct {
	Request http.Request
}
type Controller func (c Context) (interface{}, int)
