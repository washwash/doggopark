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


func ProductList(c api.Context) (interface{}, int) {
	var data [1]interface{}
	product := Product{Id: 1, Title: "name"}
	productData := make(map[string]string)
	productData["title"] = product.Title
	productData["id"] = strconv.Itoa(product.Id)
	data[0] = productData
	return data, 200

}
