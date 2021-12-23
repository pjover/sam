package model

type Consumption struct {
	Code            string
	ChildCode       int
	ProductID       string
	Units           float64
	YearMonth       string
	IsRectification bool
	InvoiceCode     string
}
