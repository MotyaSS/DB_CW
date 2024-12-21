package entity

import (
	"github.com/shopspring/decimal"
)

type Instrument struct {
	InstrumentId   int             `json:"instrument_id" db:"instrument_id"`
	CategoryId     int             `json:"category_id" db:"category_id"`
	StoreId        int             `json:"store_id" db:"store_id"`
	ManufacturerId int             `json:"manufacturer_id" db:"manufacturer_id"`
	InstrumentName string          `json:"instrument_name" db:"instrument_name"`
	Description    string          `json:"description" db:"description"`
	PricePerDay    decimal.Decimal `json:"price_per_day" db:"price_per_day"`
}

type Manufacturer struct {
	ManufacturerId   int    `json:"manufacturer_id"`
	ManufacturerName string `json:"manufacturer_name"`
}

type Category struct {
	CategoryId          int    `json:"category_id"`
	CategoryDescription string `json:"category_description"`
	CategoryName        string `json:"category_name"`
}
