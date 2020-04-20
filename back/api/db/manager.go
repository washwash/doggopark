package db

import (
  "fmt"
  "encoding/json"
)


type IModel interface{}
type QueryManager struct {
	Table   string
	Model   IModel
}

func (this QueryManager) All() []map[string]interface{} {
	query := fmt.Sprintf("SELECT * from %s;", this.Table)
	rows := Select(query)
	return rows
}


func (this QueryManager) FillStructWithRaw (raw *map[string]interface{}, model IModel) {
	bData, _ := json.Marshal(raw)
	json.Unmarshal(bData, model)
}

