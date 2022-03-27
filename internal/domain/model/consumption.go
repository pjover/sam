package model

const ConsumptionIdLength = 8

type Consumption struct {
	Id              string
	ChildId         int
	ProductId       string
	Units           float64
	YearMonth       YearMonth
	Note            string
	IsRectification bool
	InvoiceId       string
}
