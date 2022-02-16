package model

const ConsumptionCodeLength = 8

type Consumption struct {
	Code            string
	ChildCode       int
	ProductId       string
	Units           float64
	YearMonth       string
	Note            string
	IsRectification bool
	InvoiceCode     string
}
