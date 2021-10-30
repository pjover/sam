package cmd

import (
	"errors"
	"fmt"
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
		{
			args: []string{"2220", "0.5", "MME", "2", "AGE", "1", "QME", "1", "BAB"},
			expectedValue: cons.InsertConsumptionsArgs{
				Code:         2220,
				Consumptions: map[string]float64{"MME": 0.5, "AGE": 2, "QME": 1, "BAB": 1},
				Note:         testNote,
			},
			expectedError: nil,
		},
		{
			args:          []string{"2220", "0,5", "MME", "2", "AGE"},
			expectedValue: cons.InsertConsumptionsArgs{},
			expectedError: errors.New("El número introduit és invàlid: 0,5"),
		},
		{
			args:          []string{"2220", "0.5", "MME", "2"},
			expectedValue: cons.InsertConsumptionsArgs{},
			expectedError: errors.New("No s'ha indroduit el codi del darrer consum"),
		},
	}
	for i, test := range tests {
		fmt.Println("Test", i)
		var actual, err = parseInsertConsumptionsArgs(test.args, testNote)
		assert.Equal(t, test.expectedValue, actual)
		assert.Equal(t, test.expectedError, err)
	}
}
