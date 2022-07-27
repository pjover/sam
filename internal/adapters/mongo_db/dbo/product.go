package dbo

type Product struct {
	Id            string  `bson:"_id"`
	Name          string  `bson:"name"`
	ShortName     string  `bson:"shortName"`
	Price         float64 `bson:"price"`
	TaxPercentage float64 `bson:"taxPercentage"`
	IsSubsidy     bool    `bson:"isSubsidy"`
}
