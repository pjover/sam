package model

type Product struct {
	Id            string
	Name          string
	ShortName     string
	Price         float64
	TaxPercentage float64
	IsSubsidy     bool
}
