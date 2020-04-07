package store

import (
  "doggo/api"
)


func Index(c api.Context) (interface{}, int) {
	var data = make(map[string]string)
	data["store"] = "!"
	return data, 200
}

