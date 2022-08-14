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
				{InsertConsumptions, Arguments{"1480", "1", "QME", "2", "MME", "1", "AGE"}},
				{InsertConsumptions, Arguments{"1481", "1", "QME", "2", "MME", "1", "AGE"}},
			},
			[]string{
				"Laura Llull (1480): 467.00 €\n" +
					"  [2022-08]   1.0 x AGE :   12.00\n" +
					"  [2022-08]   2.0 x MME :  200.00\n" +
					"  [2022-08]   1.0 x QME :  255.00\n",
				"Aina Llull (1481): 467.00 €\n" +
					"  [2022-08]   1.0 x AGE :   12.00\n" +
					"  [2022-08]   2.0 x MME :  200.00\n" +
					"  [2022-08]   1.0 x QME :  255.00\n",
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
