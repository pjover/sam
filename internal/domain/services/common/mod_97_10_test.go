package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIban_CalculateControlCode(t *testing.T) {
	tests := []struct {
		name string
		iban Mod9710
		want string
	}{
		{
			name: "case 1",
			iban: NewMod9710("3066 8859 9782 5852 9057ES"),
			want: "28",
		},
		{
			name: "case 2",
			iban: NewMod9710("3001 2859 0880 2660 6142ES"),
			want: "02",
		},
		{
			name: "case 3",
			iban: NewMod9710("31795040719243310258ES"),
			want: "87",
		},
		{
			name: "case 4",
			iban: NewMod9710("3118-2176-0723-9984-7410ES"),
			want: "60",
		},
		{
			name: "case 5",
			iban: NewMod9710("HOBB", "20180707204308000"),
			want: "24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.iban.Checksum()
			assert.Equal(t, tt.want, got)
		})
	}
}
