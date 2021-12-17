package model

import "time"

type Invoice struct {
	Id            int
	CustomerID    int       `json:"customerId"`
	Date          time.Time `json:"date"`
	YearMonth     string    `json:"yearMonth"`
	ChildrenCodes []int     `json:"childrenCodes"`
	Lines         []Line    `json:"lines"`
	PaymentType   string    `json:"paymentType"`
	Note          string    `json:"note"`
	Emailed       bool      `json:"emailed"`
	Printed       bool      `json:"printed"`
	SentToBank    bool      `json:"sentToBank"`
	Links         struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Invoice struct {
			Href string `json:"href"`
		} `json:"invoice"`
	} `json:"_links"`
}

type Line struct {
	ProductID     string  `json:"productId"`
	Units         float64 `json:"units"`
	ProductPrice  float64 `json:"productPrice"`
	TaxPercentage float64 `json:"taxPercentage"`
	ChildCode     int     `json:"childCode"`
}
