package bdd

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func readTextFileIntoString(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func Test_stringBddBuilder_Build(t *testing.T) {
	tests := []struct {
		name string
		bdd  Bdd
		want string
	}{
		{
			name: "given a BDD should build the XML",
			bdd:  testBdd,
			want: readTextFileIntoString("string_bdd_builder.q1x"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stringBddBuilder{}
			gotContent := s.Build(tt.bdd)
			assert.Equal(t, tt.want, gotContent)
		})
	}
}
