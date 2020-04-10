package main

import (
  "log"

  "doggo/api"
  "doggo/api/db"
  "doggo/store"
)


func main () {
	log.Printf("server is running.")

	api.AddRoute("/store/", store.Index)
	api.AddRoute("/store/products/", store.ProductList)

	db.Migrate()
	api.Run()
}

