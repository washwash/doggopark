package store

import (
  "strconv"

  "doggo/api"
  _"doggo/api/db"
)


func Index(c api.Context) (interface{}, int) {
	var data = make(map[string]string)
	data["store"] = "0"
	return data, 200
}


type Product struct {
	id int
	name string
}


func ProductList(c api.Context) (interface{}, int) {
	var data [1]interface{}
	product := Product{id: 1, name: "name"}
	productData := make(map[string]string)
	productData["name"] = product.name
	productData["id"] = strconv.Itoa(product.id)
	data[0] = productData
	return data, 200

}
