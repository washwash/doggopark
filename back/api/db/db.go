package db

import (
  "os"
  "database/sql"
 _"github.com/lib/pq"
)


type client struct {
	db *sql.DB
}


func Open() *client {
	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	return &client{db: db}
}


func (this *client) Close() {
	this.db.Close()
}


func (this *client) Select(query string) []map[string]interface{} {
	rows, _ := this.db.Query(query)
	return this.RowsToMap(rows)
}


func (this *client) RowsToMap(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	var result []map[string]interface{}

	for rows.Next() {
		cols := make([]interface{}, len(columns))
		tempPointers := make([]interface{}, len(cols))
		for i := range columns {
			tempPointers[i] = &cols[i]
		}

		rows.Scan(tempPointers...)

		item := make(map[string]interface{})
		for i, name := range columns {
			item[name] = cols[i]
		}
		result = append(result, item)
	}
	return result
}

