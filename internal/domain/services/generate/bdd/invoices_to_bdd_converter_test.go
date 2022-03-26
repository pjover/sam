package bdd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_invoicesToBddConverter_calculateControlCode(t *testing.T) {
	tests := []struct {
		name   string
		params []string
		want   string
	}{
		{
			name:   "case 1",
			params: []string{"3066 8859 9782 5852 9057ES"},
			want:   "28",
		},
		{
			name:   "case 2",
			params: []string{"3001 2859 0880 2660 6142ES"},
			want:   "02",
		},
		{
			name:   "case 3",
			params: []string{"31795040719243310258ES"},
			want:   "87",
		},
		{
			name:   "case 4",
			params: []string{"3118-2176-0723-9984-7410ES"},
			want:   "60",
		},
		{
			name:   "case 5",
			params: []string{"HOBB", "20180707204308000"},
			want:   "24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := invoicesToBddConverter{}
			got := sut.calculateControlCode(tt.params...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_invoicesToBddConverter_getSepaIndentifier(t *testing.T) {
	type args struct {
		taxID   string
		country string
		suffix  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				taxID:   "36361882D",
				country: "ES",
				suffix:  "000",
			},
			want: "ES4200036361882D",
		},
		{
			name: "Case 2",
			args: args{
				taxID:   "37866397W",
				country: "ES",
				suffix:  "000",
			},
			want: "ES5500037866397W",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := invoicesToBddConverter{}
			got := sut.getSepaIndentifier(tt.args.taxID, tt.args.country, tt.args.suffix)
			assert.Equal(t, tt.want, got)
		})
	}
}
