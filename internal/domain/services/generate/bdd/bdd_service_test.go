package bdd

import (
	"testing"
)

func Test_bddService_getNextBddFilename(t *testing.T) {
	type args struct {
		currentFilenames []string
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
		}, {
			"Should return the 2nd file if there are no files",
			args{[]string{"bdd-1.qx1"}},
			"bdd-2.qx1",
		}, {
			"Should return the 4th file if there are no files",
			args{[]string{"bdd-1.qx1", "bdd-2.qx1", "bdd-3.qx1"}},
			"bdd-4.qx1",
		}, {
			"Should detect duplicates",
			args{[]string{"bdd-1.qx1", "bdd-3.qx1", "bdd-4.qx1"}},
			"bdd-5.qx1",
		}, {
			"Should detect duplicates with gap",
			args{[]string{"bdd-1.qx1", "bdd-4.qx1", "bdd-5.qx1"}},
			"bdd-6.qx1",
		}, {
			"Should detect duplicates with gap and unordered list",
			args{[]string{"bdd-7.qx1", "bdd-4.qx1", "bdd-5.qx1"}},
			"bdd-6.qx1",
		},
	}

	sut := bddService{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sut.getNextBddFilename(tt.args.currentFilenames); got != tt.want {
				t.Errorf("getNextBddFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
