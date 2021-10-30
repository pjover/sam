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
	}
	for _, test := range tests {
		var actual, err = parseInsertConsumptionsArgs(test.args, testNote)
		assert.Equal(t, test.expected, actual)
		assert.Nil(t, err)
	}
}
