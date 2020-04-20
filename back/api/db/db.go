package db

import (
  "os"
  "database/sql"
 _"github.com/lib/pq"
)


type client struct {
	db *sql.DB
}


func open() *client {
	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	return &client{db: db}
}


func (this *client) close() {
	this.db.Close()
}


func (this *client) query(q string) *sql.Rows {
	rows, _ := this.db.Query(q)
	return rows
}


func Select(query string) []map[string]interface{} {
	rows := Query(query)
	return RowsToMap(rows)
}


func Query(query string) *sql.Rows {
	conn := open()
	defer conn.close()

	return conn.query(query)
}


func RowsToMap(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	var result []map[string]interface{}

	for rows.Next() {
		tuples := make([]interface{}, len(columns))
		tempPointers := make([]interface{}, len(tuples))
		for i := range columns {
			tempPointers[i] = &tuples[i]
		}
		rows.Scan(tempPointers...)

		item := make(map[string]interface{})
		for i, name := range columns {
			var value interface{}
			val := tuples[i]
			b, ok := val.([]byte)
			if ok {
				value = string(b)
			} else {
				value = val
			}
			item[name] = value
		}
		result = append(result, item)
	}
	return result
}

