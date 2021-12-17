package dbo

type Product struct {
	Id            int     `bson:"_id"`
	Name          string  `bson:"name"`
	ShortName     string  `bson:"shortName"`
	Price         float64 `bson:"price"`
	TaxPercentage float64 `bson:"taxPercentage"`
	IsSubsidy     bool    `bson:"isSubsidy"`
}
