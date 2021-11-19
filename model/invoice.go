package model

type Line struct {
	ProductID     string  `json:"productId"`
	Units         float32 `json:"units"`
	ProductPrice  float32 `json:"productPrice"`
	TaxPercentage float32 `json:"taxPercentage"`
	ChildCode     int     `json:"childCode"`
}

type Invoice struct {
	CustomerID    int    `json:"customerId"`
	Date          string `json:"date"`
	YearMonth     string `json:"yearMonth"`
	ChildrenCodes []int  `json:"childrenCodes"`
	Lines         []Line `json:"lines"`
	PaymentType   string `json:"paymentType"`
	Note          string `json:"note"`
	Emailed       bool   `json:"emailed"`
	Printed       bool   `json:"printed"`
	SentToBank    bool   `json:"sentToBank"`
	Links         struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Invoice struct {
			Href string `json:"href"`
		} `json:"invoice"`
	} `json:"_links"`
}
