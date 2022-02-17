package billing

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testNote = "Test note"

func Test_parseInsertConsumptionsArgs(t *testing.T) {
	type args struct {
		args    []string
		noteArg string
	}
	type want struct {
		id           int
		consumptions map[string]float64
		note         string
		err          error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"Should fail with invalid child id",
			args{[]string{"0.5", "MME", "2", "MME"}, testNote},
			want{err: errors.New("el codi d'infant introduit és invàlid: 0.5")},
		}, {
			"Should fail with invalid unit format",
			args{[]string{"2220", "0,5", "MME", "2", "AGE"}, testNote},
			want{err: errors.New("el número introduit és invàlid: 0,5")},
		}, {
			"Should fail with invalid unit format",
			args{[]string{"2220", "MME", "2"}, testNote},
			want{err: errors.New("el número introduit és invàlid: MME")},
		}, {
			"Should fail with invalid product id",
			args{[]string{"2220", "2", "MMME"}, testNote},
			want{err: errors.New("el codi de producte introduit és invàlid: MMME")},
		}, {
			"Should fail without last product id",
			args{[]string{"2220", "0.5", "MME", "2"}, testNote},
			want{err: errors.New("no s'ha indroduit el codi del darrer producte")},
		}, {
			"Should fail with duplicated product id",
			args{[]string{"2220", "0.5", "MME", "2", "MME"}, testNote},
			want{err: errors.New("hi ha un codi de producte repetit")},
		}, {
			"Should parse with 1 consumption",
			args{[]string{"1552", "1", "QME"}, testNote},
			want{id: 1552, consumptions: map[string]float64{"QME": 1}, note: testNote, err: nil},
		}, {
			"Should parse with 2 consum",
			args{[]string{"2220", "0.5", "MME", "2", "AGE"}, testNote},
			want{id: 2220, consumptions: map[string]float64{"MME": 0.5, "AGE": 2}, note: testNote, err: nil},
		}, {
			"Should parse with 3 consum",
			args{[]string{"2220", "0.5", "MME", "2", "AGE", "1", "QME", "1", "BAB"}, testNote},
			want{id: 2220, consumptions: map[string]float64{"MME": 0.5, "AGE": 2, "QME": 1, "BAB": 1}, note: testNote, err: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, con, note, err := parseInsertConsumptionsArgs(tt.args.args, tt.args.noteArg)
			assert.Equal(t, tt.want.id, id)
			assert.Equal(t, tt.want.consumptions, con)
			assert.Equal(t, tt.want.note, note)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
