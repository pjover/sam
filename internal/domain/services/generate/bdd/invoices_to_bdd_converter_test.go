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
