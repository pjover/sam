package e2e

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_e2eMain(t *testing.T) {
	tests := []struct {
		name     string
		commands []Command
		want     []string
	}{
		{
			"Usual complete monthly cycle",
			[]Command{
				{InsertConsumptions, Arguments{"1480", "1", "TST", "2", "XXX", "1", "YYY"}},
				{InsertConsumptions, Arguments{"1481", "1", "TST", "3", "XXX", "0.5", "YYY"}},
				{InsertConsumptions, Arguments{"1490", "2", "TST", "5", "XXX"}},
				{InsertConsumptions, Arguments{"1491", "2", "TST", "5", "YYY"}},
				{ListConsumptions, Arguments{}},
			},
			[]string{
				"Laura Llull (1480): 34.10 €\n" +
					"  [2022-08]   1.0 x TST :   10.90\n" +
					"  [2022-08]   2.0 x XXX :   18.20\n" +
					"  [2022-08]   1.0 x YYY :    5.00\n",
				"Aina Llull (1481): 40.70 €\n" +
					"  [2022-08]   1.0 x TST :   10.90\n" +
					"  [2022-08]   3.0 x XXX :   27.30\n" +
					"  [2022-08]   0.5 x YYY :    2.50\n",
				"Antònia Petit (1490): 67.30 €\n" +
					"  [2022-08]   2.0 x TST :   21.80\n" +
					"  [2022-08]   5.0 x XXX :   45.50\n",
				"Antoni Petit (1491): 46.80 €\n" +
					"  [2022-08]   2.0 x TST :   21.80\n" +
					"  [2022-08]   5.0 x YYY :   25.00\n",
				"  Laura Llull (1480): 34.10 €\n" +
					"    [2022-08]   1.0 x TST :   10.90\n" +
					"    [2022-08]   2.0 x XXX :   18.20\n" +
					"    [2022-08]   1.0 x YYY :    5.00\n" +
					"  Aina Llull (1481): 40.70 €\n" +
					"    [2022-08]   1.0 x TST :   10.90\n" +
					"    [2022-08]   3.0 x XXX :   27.30\n" +
					"    [2022-08]   0.5 x YYY :    2.50\n" +
					"\n" +
					"2 infant(s) amb Rebut: 74.80 €\n" +
					"\n" +
					"  Antònia Petit (1490): 67.30 €\n" +
					"    [2022-08]   2.0 x TST :   21.80\n" +
					"    [2022-08]   5.0 x XXX :   45.50\n" +
					"  Antoni Petit (1491): 46.80 €\n" +
					"    [2022-08]   2.0 x TST :   21.80\n" +
					"    [2022-08]   5.0 x YYY :   25.00\n" +
					"\n" +
					"2 infant(s) amb Tranferència: 114.10 €\n" +
					"\n" +
					"TOTAL: 188.90 €",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := e2eMain(tt.commands)
			assert.Equal(t, tt.want, got)
		})
	}
}
