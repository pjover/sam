package group_type

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGroupType(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  GroupType
	}{
		{
			name:  "Undefined",
			value: "anything",
			want:  Undefined,
		},
		{
			name:  "EI_1",
			value: "EI_1",
			want:  EI_1,
		},
		{
			name:  "EI_2",
			value: "eI_2",
			want:  EI_2,
		},
		{
			name:  "EI_3",
			value: "EI_3",
			want:  EI_3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGroupType(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGroupType_Format(t *testing.T) {
	tests := []struct {
		name string
		g    GroupType
		want string
	}{
		{
			name: "Undefined",
			g:    Undefined,
			want: "Indefinit",
		},
		{
			name: "EI_1",
			g:    EI_1,
			want: "EI 1 (0-1)",
		},
		{
			name: "EI_2",
			g:    EI_2,
			want: "EI 2 (1-2)",
		},
		{
			name: "EI_3",
			g:    EI_3,
			want: "EI 3 (2-3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Format(); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
