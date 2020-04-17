package store

import (
  "doggo/api"
)


func Index(c api.Context) (interface{}, int) {
	var data = make(map[string]string)
	return data, 200
}


func ProductList(c api.Context) (interface{}, int) {
	manager := GetProductManager()

	raw := manager.All()
	products := manager.RawToModel(raw)
	result := products
	return result, 200
}
