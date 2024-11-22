package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Rental struct {
	RentalId     int       `json:"rental_id"`
	UserId       int       `json:"user_id"`
	InstrumentId int       `json:"instrument_id"`
	RentalDate   time.Time `json:"rental_date"`
	ReturnDate   time.Time `json:"return_date"`
}

type Payment struct {
	PaymentId     int             `json:"payment_id"`
	RentalId      int             `json:"rental_id"`
	PaymentDate   time.Time       `json:"payment_date"`
	PaymentAmount decimal.Decimal `json:"payment_amount"`
}
