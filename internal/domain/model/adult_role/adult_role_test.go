package adult_role

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAdultRole(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  AdultRole
	}{
		{
			name:  "Invalid",
			value: "Anything",
			want:  Invalid,
		},
		{
			name:  "Mother",
			value: "MOTHER",
			want:  Mother,
		},
		{
			name:  "Father",
			value: "FATHER",
			want:  Father,
		},
		{
			name:  "Tutor",
			value: "TUTOR",
			want:  Tutor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAdultRole(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAdultRole_Format(t *testing.T) {
	tests := []struct {
		name string
		p    AdultRole
		want string
	}{
		{
			name: "Indefinit",
			p:    Invalid,
			want: "Indefinit",
		},
		{
			name: "Mare",
			p:    Mother,
			want: "Mare",
		},
		{
			name: "Pare",
			p:    Father,
			want: "Pare",
		},
		{
			name: "Tutor",
			p:    Tutor,
			want: "Tutor",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Format()
			assert.Equal(t, tt.want, got)
		})
	}
}
