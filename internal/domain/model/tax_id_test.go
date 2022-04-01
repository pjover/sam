package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTaxId(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		{
			name:    "empty TaxiId",
			id:      "",
			wantErr: nil,
		},
		{
			name:    "correct CIF 1",
			id:      "F9873859D",
			wantErr: nil,
		},
		{
			name:    "correct CIF 2",
			id:      "H69159929",
			wantErr: nil,
		},
		{
			name:    "correct CIF 3",
			id:      "N4585129B",
			wantErr: nil,
		},
		{
			name:    "correct CIF 4",
			id:      "J8518436D",
			wantErr: nil,
		},
		{
			name:    "correct CIF 5",
			id:      "S2910252B",
			wantErr: nil,
		},
		{
			name:    "correct CIF 6",
			id:      "R2305254A",
			wantErr: nil,
		},
		{
			name:    "correct CIF 7",
			id:      "A58818501",
			wantErr: nil,
		},
		{
			name:    "correct CIF 8",
			id:      "Z4614782K",
			wantErr: nil,
		},
		{
			name:    "correct CIF 9",
			id:      "Q8453568A",
			wantErr: nil,
		},
		{
			name:    "incorrect CIF 1",
			id:      "N8486733K",
			wantErr: errors.New("el DNI/NIE/CIF 'N8486733K' no és vàlid"),
		},
		{
			name:    "incorrect CIF 2",
			id:      "J8528436D",
			wantErr: errors.New("el DNI/NIE/CIF 'J8528436D' no és vàlid"),
		},
		{
			name:    "incorrect CIF 3",
			id:      "F9883859D",
			wantErr: errors.New("el DNI/NIE/CIF 'F9883859D' no és vàlid"),
		},
		{
			name:    "incorrect CIF 4",
			id:      "H69169929",
			wantErr: errors.New("el DNI/NIE/CIF 'H69169929' no és vàlid"),
		},
		{
			name:    "incorrect CIF 5",
			id:      "N4586129B",
			wantErr: errors.New("el DNI/NIE/CIF 'N4586129B' no és vàlid"),
		},
		{
			name:    "incorrect CIF 6",
			id:      "S2919252B",
			wantErr: errors.New("el DNI/NIE/CIF 'S2919252B' no és vàlid"),
		},
		{
			name:    "incorrect CIF 7",
			id:      "R2306254A",
			wantErr: errors.New("el DNI/NIE/CIF 'R2306254A' no és vàlid"),
		},
		{
			name:    "incorrect CIF 8",
			id:      "A58838501",
			wantErr: errors.New("el DNI/NIE/CIF 'A58838501' no és vàlid"),
		},
		{
			name:    "incorrect CIF 9",
			id:      "Y58838501",
			wantErr: errors.New("el DNI/NIE/CIF 'Y58838501' no és vàlid"),
		},
		{
			name:    "correct NIE 1",
			id:      "X2010159M",
			wantErr: nil,
		},
		{
			name:    "incorrect NIE 1",
			id:      "X2010159K",
			wantErr: errors.New("el DNI/NIE/CIF 'X2010159K' no és vàlid"),
		},
		{
			name:    "incorrect NIE 2",
			id:      "Y2010159K",
			wantErr: errors.New("el DNI/NIE/CIF 'Y2010159K' no és vàlid"),
		},
		{
			name:    "incorrect NIE 3",
			id:      "Z2010159K",
			wantErr: errors.New("el DNI/NIE/CIF 'Z2010159K' no és vàlid"),
		},
		{
			name:    "incorrect NIE 4",
			id:      "Z2010159",
			wantErr: errors.New("el DNI/NIE/CIF 'Z2010159' no és vàlid"),
		},
		{
			name:    "correct NIF 1",
			id:      "48592145E",
			wantErr: nil,
		},
		{
			name:    "incorrect NIF 1",
			id:      "48592145K",
			wantErr: errors.New("el DNI/NIE/CIF '48592145K' no és vàlid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTaxId(tt.id)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
