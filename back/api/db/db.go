package db

import (
  "os"
  "log"
  "fmt"

  "database/sql"
  _"github.com/lib/pq"
  "github.com/golang-migrate/migrate/v4"
  _"github.com/golang-migrate/migrate/v4/database/postgres"
  _"github.com/golang-migrate/migrate/v4/source/file"
)


type DataBaseClient interface{
}
type client struct {
	db *sql.DB
}


var (
	Client DataBaseClient
)


func init() {
	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	Client = &client{db: db}
}


func Migrate() {
	log.Printf("Migration start...")
	defer log.Printf("Migration is complete!")

	migrationsPath := fmt.Sprintf("file:/%s", os.Getenv("DB_MIGRATIONS_PATH"))
	m, err := migrate.New(migrationsPath, os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}

