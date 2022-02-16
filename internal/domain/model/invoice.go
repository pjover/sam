package model

import "time"

type Invoice struct {
	Code          string
	CustomerID    int
	Date          time.Time
	YearMonth     string
	ChildrenCodes []int
	Lines         []Line
	PaymentType   string
	Note          string
	Emailed       bool
	Printed       bool
	SentToBank    bool
}

type Line struct {
	ProductId     string
	Units         float64
	ProductPrice  float64
	TaxPercentage float64
	ChildCode     int
}
