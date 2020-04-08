package db

import (
  "os"
  "fmt"
  "database/sql"
)


type DataBaseClient interface{}
type client struct {
	db *sql.DB
}


var (
	Client DataBaseClient
)


func init() {
	user, pwd := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	conn := fmt.Sprintf("postgres://%s:%s@localhost/%s", user, pwd, dbName)
	db, _ := sql.Open("postgres", conn)
	Client = &client{db: db}
}

