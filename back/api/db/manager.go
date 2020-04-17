package db

import (
  _"log"
  "reflect"
  "fmt"
  "encoding/json"
)


type IModel interface{}
type FieldPair struct {
	field  reflect.StructField
	alias  string
}

type IQueryManager interface{
	All() []map[string]string
	Filter(map[string]string) []IModel

	Get(map[string]string) IModel
	Create(map[string]string) IModel
	Update(map[string]string) IModel
	Delete(map[string]string)
}

type QueryManager struct {
	Table   string
	Model   IModel

	fields  []FieldPair
}


func (this QueryManager) All() []map[string]interface{} {
	conn := Open()
	defer conn.Close()

	query := fmt.Sprintf("SELECT * from %s;", this.Table)
	rows := conn.Select(query)
	return rows
}


func (this QueryManager) FillStructWithRaw (raw *map[string]interface{}, model IModel) {
	bData, _ := json.Marshal(raw)
	json.Unmarshal(bData, model)
}

/*
{
	t := reflect.TypeOf(this.Model)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			continue
		}
		this.fields = append(this.fields, FieldPair{field, tag})
	}
}
*/
