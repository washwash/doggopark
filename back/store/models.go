package store

import (
  d "github.com/shopspring/decimal"

  "doggo/api/db"
)


type price struct {
	Id            int
	Amount        d.Decimal
	CurrencyCode  string
}


type product struct {
	Id            int
	ExternalId    string
	Title         string
	Description   string
	Price         price

}


type productManager struct {
	db.QueryManager
}


func (this productManager) RawToModel(raw []map[string]interface{}) []product {
	result := make([]product, 0, len(raw))
	for _, r := range raw {
		p := new(product)
		this.FillStructWithRaw(&r, p)
		result = append(result, *p)
	}
	return result
}


func GetProductManager () productManager {
	baseManager := db.QueryManager{Table: "product", Model: product{}}
	return productManager{baseManager}
}

