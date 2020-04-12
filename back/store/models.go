package store

import (
  d "github.com/shopspring/decimal"

  "doggo/api/db"
)


type Price struct {
	db.Model

	Id            int
	Amount        d.Decimal
	CurrencyCode  string
}


type Product struct {
	db.Model

	Id           int
	ExternalId   string
	Title        string
	Description  string
	PriceId      int

	Price        Price
}

