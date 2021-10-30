package cons

type InsertConsumptionsArgs struct {
	Code         int
	Consumptions map[string]float64
	Note         string
}

func InsertConsumptions(args InsertConsumptionsArgs) error {
	return nil
}
