package store

import (
  "doggo/api"
)


func Index(w api.IResponseWriter, r *api.IRequest) {
	w.Header().Set("key", "value")
	w.Write([]byte("store"))
}

