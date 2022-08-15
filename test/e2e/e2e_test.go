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
				{BillConsumptions, Arguments{}},
			},
			[]string{
				"Laura Llull (1480): 34.60 €\n" +
					"  [2022-08]   1.0 x TST :   10.90\n" +
					"  [2022-08]   2.0 x XXX :   18.20\n" +
					"  [2022-08]   1.0 x YYY :    5.50\n",
				"Aina Llull (1481): 40.95 €\n" +
					"  [2022-08]   1.0 x TST :   10.90\n" +
					"  [2022-08]   3.0 x XXX :   27.30\n" +
					"  [2022-08]   0.5 x YYY :    2.75\n",
				"Antònia Petit (1490): 67.30 €\n" +
					"  [2022-08]   2.0 x TST :   21.80\n" +
					"  [2022-08]   5.0 x XXX :   45.50\n",
				"Antoni Petit (1491): 49.30 €\n" +
					"  [2022-08]   2.0 x TST :   21.80\n" +
					"  [2022-08]   5.0 x YYY :   27.50\n",
				"  Laura Llull (1480): 34.60 €\n" +
					"    [2022-08]   1.0 x TST :   10.90\n" +
					"    [2022-08]   2.0 x XXX :   18.20\n" +
					"    [2022-08]   1.0 x YYY :    5.50\n" +
					"  Aina Llull (1481): 40.95 €\n" +
					"    [2022-08]   1.0 x TST :   10.90\n" +
					"    [2022-08]   3.0 x XXX :   27.30\n" +
					"    [2022-08]   0.5 x YYY :    2.75\n" +
					"\n" +
					"2 infant(s) amb Rebut: 75.55 €\n" +
					"\n" +
					"  Antònia Petit (1490): 67.30 €\n" +
					"    [2022-08]   2.0 x TST :   21.80\n" +
					"    [2022-08]   5.0 x XXX :   45.50\n" +
					"  Antoni Petit (1491): 49.30 €\n" +
					"    [2022-08]   2.0 x TST :   21.80\n" +
					"    [2022-08]   5.0 x YYY :   27.50\n" +
					"\n" +
					"2 infant(s) amb Tranferència: 116.60 €\n" +
					"\n" +
					"TOTAL: 192.15 €",
				" 1. Cara Santamaria 148  F-34  2022-08    75.55  Rebut  0.5 YYY (2.75),1.0 TST (10.90),1.0 TST (10.90),1.0 YYY (5.50),2.0 XXX (18.20),3.0 XXX (27.30)\n" +
					"Total 1 Rebut: 75.55 €\n" +
					" 1. Joana Petita 149  X-23  2022-08   116.60  Tranferència  2.0 TST (21.80),2.0 TST (21.80),5.0 XXX (45.50),5.0 YYY (27.50)\n" +
					"Total 1 Tranferència: 116.60 €\n" +
					"TOTAL: 192.15 €\n",
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
