package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsumptionListToString(t *testing.T) {
	type args struct {
		consumptions []Consumption
		child        Child
		products     map[string]Product
		indentText   string
	}
	tests := []struct {
		name      string
		args      args
		wantText  string
		wantTotal float64
	}{
		{
			name: "ConsumptionListToString",
			args: args{
				consumptions: []Consumption{
					{
						childId:   1480,
						productId: "TST",
						units:     2,
						yearMonth: YearMonth{2022, 7},
						note:      "Note 1",
					},
					{
						childId:   1480,
						productId: "XXX",
						units:     1,
						yearMonth: YearMonth{2022, 7},
						note:      "Note 2",
					},
				},
				child: TestChild1480,
				products: map[string]Product{
					"TST": ProductTST,
					"XXX": ProductXXX,
				},
				indentText: "",
			},
			wantText: "Laura Llull (1480): 30.90 €\n" +
				"  [2022-07]   2.0 x TST :   21.80\n" +
				"  [2022-07]   1.0 x XXX :    9.10\n",
			wantTotal: 30.9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotText, gotTotal := ConsumptionListFormatValues(tt.args.consumptions, tt.args.child, tt.args.products, tt.args.indentText)
			assert.Equal(t, tt.wantText, gotText)
			assert.Equal(t, tt.wantTotal, gotTotal)
		})
	}
}
