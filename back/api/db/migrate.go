package db

import (
  "os"
  "log"
  "fmt"

  "github.com/golang-migrate/migrate/v4"
  _"github.com/golang-migrate/migrate/v4/database/postgres"
  _"github.com/golang-migrate/migrate/v4/source/file"

)


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

