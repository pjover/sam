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

var testNote = "Test note"

func TestParseInsertConsumptionsArgs(t *testing.T) {

	tests := []testData{
		{
			args: []string{"1552", "1", "QME"},
			expected: cons.InsertConsumptionsArgs{
				Code:         1552,
				Consumptions: map[string]float64{"QME": 1},
				Note:         testNote,
			},
		},
		{
			args: []string{"2220", "0.5", "MME", "2", "AGE"},
			expected: cons.InsertConsumptionsArgs{
				Code:         2220,
				Consumptions: map[string]float64{"MME": 0.5, "AGE": 2},
				Note:         testNote,
			},
		},
	}
	for _, test := range tests {
		var actual, err = parseInsertConsumptionsArgs(test.args, testNote)
		assert.Equal(t, test.expected, actual)
		assert.Nil(t, err)
	}
}
