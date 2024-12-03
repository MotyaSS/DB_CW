package entity

import (
	"github.com/shopspring/decimal"
)

type InstFilter struct {
	Category     *string
	Manufacturer *string
	PriceFloor   *decimal.Decimal
	PriceCeil    *decimal.Decimal
	Page         int
}

func (f *InstFilter) AddCategory(category string) *InstFilter {
	f.Category = &category
	return f
}

func (f *InstFilter) AddManufacturer(manufacturer string) *InstFilter {
	f.Manufacturer = &manufacturer
	return f
}

func (f *InstFilter) AddPriceFloor(price *decimal.Decimal) *InstFilter {
	f.PriceFloor = price
	return f
}

func (f *InstFilter) AddPriceCeil(price *decimal.Decimal) *InstFilter {
	f.PriceCeil = price
	return f
}

func (f *InstFilter) AddPage(page int) *InstFilter {
	f.Page = page
	return f
}
