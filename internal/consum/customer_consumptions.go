package consum

type CustomerConsumptionsManager interface {
	Run(args CustomerConsumptionsArgs) (string, error)
}

type CustomerConsumptionsArgs struct {
	Code         int
	Consumptions map[string]float64
	Note         string
}
