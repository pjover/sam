package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extractCountryCode(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    string
		wantErr error
	}{
		{
			name: "All strings",
			code: "ES2830668859978258529057",
			want: "ES",
		},
		{
			name: "Upper case",
			code: "es2830668859978258529057",
			want: "ES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractCountryCode(tt.code)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
