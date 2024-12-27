package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Discount struct {
	DiscountId         int             `json:"discount_id" db:"discount_id"`
	InstrumentId       int             `json:"instrument_id" db:"instrument_id"`
	DiscountPercentage decimal.Decimal `json:"discount_percentage" db:"discount_percentage"`
	ValidUntil         time.Time       `json:"valid_until" db:"valid_until"`
}
