package cmd

import (
	"errors"
	"reflect"
	"sam/core"
	"testing"
)

var testNote = "Test note"

func Test_parseInsertConsumptionsArgs(t *testing.T) {
	type args struct {
		args    []string
		noteArg string
	}
	tests := []struct {
		name    string
		args    args
		want    core.InsertConsumptionsArgs
		wantErr error
	}{
		{
			"Should fail with invalid child code",
			args{[]string{"0.5", "MME", "2", "MME"}, testNote},
			core.InsertConsumptionsArgs{},
			errors.New("El codi d'infant introduit és invàlid: 0.5"),
		}, {
			"Should fail with invalid unit format",
			args{[]string{"2220", "0,5", "MME", "2", "AGE"}, testNote},
			core.InsertConsumptionsArgs{},
			errors.New("El número introduit és invàlid: 0,5"),
		}, {
			"Should fail with invalid unit format",
			args{[]string{"2220", "MME", "2"}, testNote},
			core.InsertConsumptionsArgs{},
			errors.New("El número introduit és invàlid: MME"),
		}, {
			"Should fail with invalid product code",
			args{[]string{"2220", "2", "MMME"}, testNote},
			core.InsertConsumptionsArgs{},
			errors.New("El codi de producte introduit és invàlid: MMME"),
		}, {
			"Should fail without last product code",
			args{[]string{"2220", "0.5", "MME", "2"}, testNote},
			core.InsertConsumptionsArgs{},
			errors.New("No s'ha indroduit el codi del darrer producte"),
		}, {
			"Should fail with duplicated product code",
			args{[]string{"2220", "0.5", "MME", "2", "MME"}, testNote},
			core.InsertConsumptionsArgs{},
			errors.New("Hi ha un codi de producte repetit"),
		}, {
			"Should parse with 1 consumption",
			args{[]string{"1552", "1", "QME"}, testNote},
			core.InsertConsumptionsArgs{Code: 1552, Consumptions: map[string]float64{"QME": 1}, Note: testNote},
			nil,
		}, {
			"Should parse with 2 consumptions",
			args{[]string{"2220", "0.5", "MME", "2", "AGE"}, testNote},
			core.InsertConsumptionsArgs{Code: 2220, Consumptions: map[string]float64{"MME": 0.5, "AGE": 2}, Note: testNote},
			nil,
		}, {
			"Should parse with 3 consumptions",
			args{[]string{"2220", "0.5", "MME", "2", "AGE", "1", "QME", "1", "BAB"}, testNote},
			core.InsertConsumptionsArgs{Code: 2220, Consumptions: map[string]float64{"MME": 0.5, "AGE": 2, "QME": 1, "BAB": 1}, Note: testNote},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInsertConsumptionsArgs(tt.args.args, tt.args.noteArg)
			if (err != nil) && !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("parseInsertConsumptionsArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInsertConsumptionsArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
