package util

import "testing"

func Test_extractDefaultName(t *testing.T) {
	type args struct {
		contentDisposition string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Extract filename",
			args{`form-data; name="attachment"; filename="F-3945 (227).pdf"`},
			"F-3945 (227).pdf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractDefaultName(tt.args.contentDisposition); got != tt.want {
				t.Errorf("extractDefaultName() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
