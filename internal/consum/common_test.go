package consum

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildConsumptionsBody(t *testing.T) {

	t.Run("Should return a basic ConsumptionsBody", func(t *testing.T) {
		var args = CustomerConsumptionsArgs{
			Code:         222,
			Consumptions: map[string]float64{"AAA": 1.5},
			Note:         "Test note",
		}
		actual := buildConsumptionsBody(args)

		var expected = Body{
			YearMonth: viper.GetString("yearMonth"),
			Children: []ChildBody{
				{
					Code: 222,
					Consumptions: []ConsumptionBody{
						{"AAA", 1.5, "Test note"},
					},
				},
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return a multiple ConsumptionsBody", func(t *testing.T) {
		var args = CustomerConsumptionsArgs{
			Code:         222,
			Consumptions: map[string]float64{"AAA": 1.5, "BBB": 2, "CCC": 7.7},
			Note:         "Test note",
		}
		actual := buildConsumptionsBody(args)

		var expected = Body{
			YearMonth: viper.GetString("yearMonth"),
			Children: []ChildBody{
				{
					Code: 222,
					Consumptions: []ConsumptionBody{
						{"AAA", 1.5, "Test note"},
						{"BBB", 2, ""},
						{"CCC", 7.7, ""},
					},
				},
			},
		}
		assert.Equal(t, expected, actual)
	})
}
