package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Repair struct {
	RepairId        int             `json:"repair_id"`
	InstrumentId    int             `json:"instrument_id"`
	RepairStartDate time.Time       `json:"repair_start_date"`
	RepairEndDate   time.Time       `json:"repair_end_date"`
	RepairCost      decimal.Decimal `json:"repair_cost"`
	Description     string          `json:"description"`
}
