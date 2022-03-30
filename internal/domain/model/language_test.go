package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLanguage(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  Language
	}{
		{
			name:  "Undefined",
			value: "anything",
			want:  Undefined,
		},
		{
			name:  "Catalan",
			value: "ca",
			want:  Catalan,
		},
		{
			name:  "English",
			value: "EN",
			want:  English,
		},
		{
			name:  "Spanish",
			value: "ES",
			want:  Spanish,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLanguage(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLanguage_Format(t *testing.T) {
	tests := []struct {
		name string
		lang Language
		want string
	}{
		{
			"Undefined",
			Undefined,
			"Indefinit",
		},
		{
			"Catalan",
			Catalan,
			"Català",
		},
		{
			"English",
			English,
			"Anglès",
		},
		{
			"Spanish",
			Spanish,
			"Espanyol",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lang.Format()
			assert.Equal(t, tt.want, got)
		})
	}
}
