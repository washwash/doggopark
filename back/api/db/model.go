package db


type Model struct {}

type IQueryManager interface{
	All() []Model
	Filter(map[string]string) []Model

	Get(map[string]string) Model
	Create(map[string]string) Model
	Update(map[string]string) Model
	Delete(map[string]string)
}

func (this Model) All() []Model{
	result := make([]Model, 0, 1)
	result = append(result, this)
	return result
}


