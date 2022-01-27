package generate

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/pjover/sam/internal/generate/invoices/mocks"
)

func Test_GenerateSingleInvoiceCmd(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name      string
		args      args
		mockedArg string
		want      string
		wantErr   error
	}{
		{
			"Accepts invoice code",
			args{[]string{"f-3945"}},
			"F-3945",
			"Generant la factura F-3945",
			nil,
		},
	}

	mockedGenerator := new(mocks.SingleInvoiceGenerator)

	sut := newGenerateSingleInvoiceCmd(mockedGenerator)
	buffer := bytes.NewBufferString("")
	sut.SetOut(buffer)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedGenerator.On("Generate", tt.mockedArg).Return(tt.want, tt.wantErr)

			sut.SetArgs(tt.args.args)
			_ = sut.Execute()
			out, err := ioutil.ReadAll(buffer)
			got := string(out)
			if (err != nil) && !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
