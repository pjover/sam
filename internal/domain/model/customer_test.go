package model

import (
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransientCustomer_Validate(t *testing.T) {
	type fields struct {
		Children      []TransientChild
		Adults        []Adult
		InvoiceHolder InvoiceHolder
		Note          string
		Language      language.Language
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "Ok",
			fields: fields{
				Children: []TransientChild{
					{
						Name:      "Some name",
						Surname:   "Some surname",
						BirthDate: testDate,
						Group:     group_type.Ei1,
					},
				},
				Adults: []Adult{
					TestAdultMother148,
				},
				InvoiceHolder: TestInvoiceHolder148,
				Language:      language.Catalan,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			sut := TransientCustomer{
				Children:      tt.fields.Children,
				Adults:        tt.fields.Adults,
				InvoiceHolder: tt.fields.InvoiceHolder,
				Note:          tt.fields.Note,
				Language:      tt.fields.Language,
			}
			got := sut.Validate()
			assert.Equal(t, tt.wantErr, got)
		})
	}
}
