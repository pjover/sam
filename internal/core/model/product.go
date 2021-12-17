package model

type Product struct {
	Id            int
	Name          string  `json:"name"`
	ShortName     string  `json:"shortName"`
	Price         float64 `json:"price"`
	TaxPercentage float64 `json:"taxPercentage"`
	IsSubsidy     bool    `json:"isSubsidy"`
	Links         struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Product struct {
			Href string `json:"href"`
		} `json:"product"`
	} `json:"_links"`
}
