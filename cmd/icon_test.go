package cmd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sam/cons"
	"testing"
)

type testData struct {
	args          []string
	expectedValue cons.InsertConsumptionsArgs
	expectedError error
}

var testNote = "Test note"

func TestParseInsertConsumptionsArgs(t *testing.T) {

	tests := []testData{
		{
			args:          []string{"1552"},
			expectedValue: cons.InsertConsumptionsArgs{},
			expectedError: errors.New("Introdueix més de 3 arguments, has introduit 1 arguments"),
		},
		{
			args: []string{"1552", "1", "QME"},
			expectedValue: cons.InsertConsumptionsArgs{
				Code:         1552,
				Consumptions: map[string]float64{"QME": 1},
				Note:         testNote,
			},
			expectedError: nil,
		},
		{
			args: []string{"2220", "0.5", "MME", "2", "AGE"},
			expectedValue: cons.InsertConsumptionsArgs{
				Code:         2220,
				Consumptions: map[string]float64{"MME": 0.5, "AGE": 2},
				Note:         testNote,
			},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		var actual, err = parseInsertConsumptionsArgs(test.args, testNote)
		assert.Equal(t, test.expectedValue, actual)
		assert.Equal(t, test.expectedError, err)
	}
}
