package DB_CW

import (
	"github.com/shopspring/decimal"
)

type Discount struct {
	DiscountId         int             `json:"discount_id"`
	InstrumentId       int             `json:"instrument_id"`
	DiscountPercentage decimal.Decimal `json:"discount_percentage"`
}
