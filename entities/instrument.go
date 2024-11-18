package DB_CW

import (
	"github.com/shopspring/decimal"
)

type Instrument struct {
	InstrumentId   int             `json:"instrument_id"`
	CategoryId     int             `json:"category_id"`
	StoreId        int             `json:"store_id"`
	ManufacturerId int             `json:"manufacturer_id"`
	InstrumentName string          `json:"name"`
	Description    string          `json:"description"`
	PricePerDay    decimal.Decimal `json:"price_per_day"`
}

type Manufacturer struct {
	ManufacturerId   int    `json:"manufacturer_id"`
	ManufacturerName string `json:"manufacturer_name"`
	Description      string `json:"manufacturer_description"`
}

type Category struct {
	CategoryId          int    `json:"category_id"`
	CategoryDescription string `json:"category_description"`
	CategoryName        string `json:"category_name"`
}
