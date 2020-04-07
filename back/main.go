package main

import (
  "log"

  "doggo/api"
  "doggo/store"
)


func main () {
	log.Printf("server is running.")

	api.AddRoute("/store/", store.Index)
	api.AddRoute("/store/products/", store.ProductList)
	api.Run()
}

