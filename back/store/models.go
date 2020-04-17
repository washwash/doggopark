package store

import (
  d "github.com/shopspring/decimal"

  "doggo/api/db"
)


type price struct {
	Id            int        `json:"id"`
	Amount        d.Decimal  `json:"amount"`
	CurrencyCode  string     `json:"currency_code"`
}


type product struct {
	Id           int     `json:"id"`
	ExternalId   string  `json:"external_id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	PriceId      int     `json:"price_id"`

	Price        price   `json:"-"`
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

