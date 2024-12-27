package entity

import (
	"github.com/shopspring/decimal"
)

type InstFilter struct {
	Categories    []string
	Manufacturers []string
	PriceFloor    *decimal.Decimal
	PriceCeil     *decimal.Decimal
	Page          int
}

func (f *InstFilter) AddCategory(category string) *InstFilter {
	if f.Categories == nil {
		f.Categories = make([]string, 0)
	}
	f.Categories = append(f.Categories, category)
	return f
}

func (f *InstFilter) AddManufacturer(manufacturer string) *InstFilter {
	if f.Manufacturers == nil {
		f.Manufacturers = make([]string, 0)
	}
	f.Manufacturers = append(f.Manufacturers, manufacturer)
	return f
}

func (f *InstFilter) AddPriceFloor(price decimal.Decimal) *InstFilter {
	f.PriceFloor = &price
	return f
}

func (f *InstFilter) AddPriceCeil(price decimal.Decimal) *InstFilter {
	f.PriceCeil = &price
	return f
}

func (f *InstFilter) AddPage(page int) *InstFilter {
	f.Page = page
	return f
}
