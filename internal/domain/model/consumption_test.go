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
	}
	tests := []struct {
		name string
		args args
		want string
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
			},
			want: "Laura Llull (1480): 30.90 €\n" +
				"  [2022-07]   2.0 x TST :   21.80\n" +
				"  [2022-07]   1.0 x XXX :    9.10\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConsumptionListToString(tt.args.consumptions, tt.args.child, tt.args.products)
			assert.Equal(t, tt.want, got)
		})
	}
}
