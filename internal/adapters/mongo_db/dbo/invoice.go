package dbo

import (
	"time"
)

type Invoice struct {
	Id          string    `bson:"_id"`
	CustomerID  int       `bson:"customerId"`
	Date        time.Time `bson:"date"`
	YearMonth   string    `bson:"yearMonth"`
	ChildrenIds []int     `bson:"childrenCodes"`
	Lines       []Line    `bson:"lines"`
	PaymentType string    `bson:"paymentType"`
	Note        string    `bson:"note"`
	Emailed     bool      `bson:"emailed"`
	SentToBank  bool      `bson:"sentToBank"`
}

type Line struct {
	ProductID     string  `bson:"productId"`
	Units         float64 `bson:"units"`
	ProductPrice  float64 `bson:"productPrice"`
	TaxPercentage float64 `bson:"taxPercentage"`
	ChildId       int     `bson:"childCode"`
}

func (i Invoice) GetId() interface{} {
	return i.Id
}
