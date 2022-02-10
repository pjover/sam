package model

const ConsumptionCodeLength = 8

type Consumption struct {
	Code            string
	ChildCode       int
	ProductID       string
	Units           float64
	YearMonth       string
	Note            string
	IsRectification bool
	InvoiceCode     string
}
