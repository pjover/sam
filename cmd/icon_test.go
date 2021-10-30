package cmd

import (
	"github.com/stretchr/testify/assert"
	"sam/cons"
	"testing"
)

type testData struct {
	args     []string
	expected cons.InsertConsumptionsArgs
}

func TestParseInsertConsumptionsArgs(t *testing.T) {

	tests := []testData{
		{
			args: []string{"1552", "1", "QME"},
			expected: cons.InsertConsumptionsArgs{
				Code:         1552,
				Consumptions: map[string]float64{"QME": 1},
				Note:         "",
			},
		},
		{
			args: []string{"1552", "1", "QME"},
			expected: cons.InsertConsumptionsArgs{
				Code:         1552,
				Consumptions: map[string]float64{"QME": 1},
				Note:         "",
			},
		},
	}
	for _, test := range tests {
		var actual, err = parseInsertConsumptionsArgs(test.args)
		assert.Equal(t, test.expected, actual)
		assert.NotNil(t, err)
	}
}
