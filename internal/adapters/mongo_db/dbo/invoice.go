package dbo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	Id            string    `bson:"_id"`
	CustomerID    int       `bson:"customerId"`
	Date          time.Time `bson:"date"`
	YearMonth     string    `bson:"yearMonth"`
	ChildrenCodes []int     `bson:"childrenCodes"`
	Lines         []Line    `bson:"lines"`
	PaymentType   string    `bson:"paymentType"`
	Note          string    `bson:"note"`
	Emailed       bool      `bson:"emailed"`
	Printed       bool      `bson:"printed"`
	SentToBank    bool      `bson:"sentToBank"`
}

type Line struct {
	ProductID     string               `bson:"productId"`
	Units         primitive.Decimal128 `bson:"units"`
	ProductPrice  primitive.Decimal128 `bson:"productPrice"`
	TaxPercentage primitive.Decimal128 `bson:"taxPercentage"`
	ChildCode     int                  `bson:"childCode"`
}
