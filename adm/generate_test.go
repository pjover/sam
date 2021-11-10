package adm

import "testing"

func Test_getNextFilename(t *testing.T) {
	type args struct {
		filenames []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Should return the 1st file if there are no files",
			args{[]string{}},
			"bdd-1.qx1",
		},
		{
			"Should return the 2nd file if there are no files",
			args{[]string{"bdd-1.qx1"}},
			"bdd-2.qx1",
		},
		{
			"Should return the 4th file if there are no files",
			args{[]string{"bdd-1.qx1", "bdd-2.qx1", "bdd-3.qx1"}},
			"bdd-4.qx1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextFilename(tt.args.filenames); got != tt.want {
				t.Errorf("getNextFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
