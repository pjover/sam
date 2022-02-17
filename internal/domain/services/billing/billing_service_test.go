package billing

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/ports/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const yearMonth = "2022-02"

var today = time.Date(2022, 2, 16, 20, 33, 59, 0, time.Local)

var noRectificationConsumptions = []model.Consumption{
	{
		Id:        "AA1",
		ChildId:   1850,
		ProductId: "TST",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 1",
	},
	{
		Id:        "AA2",
		ChildId:   1850,
		ProductId: "TST",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 2",
	},
	{
		Id:        "AA3",
		ChildId:   1851,
		ProductId: "TST",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 3",
	},
	{
		Id:        "AA4",
		ChildId:   1851,
		ProductId: "XXX",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 4",
	},
	{
		Id:        "AA5",
		ChildId:   1860,
		ProductId: "TST",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 5",
	},
	{
		Id:        "AA7",
		ChildId:   1860,
		ProductId: "YYY",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "",
	},
	{
		Id:        "AA8",
		ChildId:   1860,
		ProductId: "YYY",
		Units:     -2,
		YearMonth: yearMonth,
		Note:      "",
	},
}

var customer185 = model.Customer{
	Id: 185,
	InvoiceHolder: model.InvoiceHolder{
		PaymentType: payment_type.BankDirectDebit,
	},
}

var customer186 = model.Customer{
	Id: 186,
	InvoiceHolder: model.InvoiceHolder{
		PaymentType: payment_type.BankDirectDebit,
	},
}

var productTST = model.Product{
	Id:            "TST",
	Price:         10.9,
	TaxPercentage: 0.0,
}

var productXXX = model.Product{
	Id:            "XXX",
	Price:         9.1,
	TaxPercentage: 0.0,
}

var productYYY = model.Product{
	Id:            "XXX",
	Price:         9.1,
	TaxPercentage: 0.0,
}

func Test_billingService_consumptionsToInvoices(t *testing.T) {
	tests := []struct {
		name         string
		consumptions []model.Consumption
		wantInvoices []model.Invoice
	}{
		{
			name:         "without rectification",
			consumptions: noRectificationConsumptions,
			wantInvoices: []model.Invoice{
				{
					Id:          "",
					CustomerId:  185,
					Date:        today,
					YearMonth:   yearMonth,
					ChildrenIds: []int{1850, 1851},
					Lines: []model.Line{
						{
							ProductId:     "TST",
							Units:         4,
							ProductPrice:  10.9,
							TaxPercentage: 0,
							ChildId:       1850,
						},
						{
							ProductId:    "TST",
							Units:        2,
							ProductPrice: 10.9,
							ChildId:      1851,
						},
						{
							ProductId:    "XXX",
							Units:        2,
							ProductPrice: 9.1,
							ChildId:      1851,
						},
					},
					PaymentType: payment_type.BankDirectDebit,
					Note:        "Note 1, Note 2, Note 3, Note 4",
				},
				{
					Id:          "",
					CustomerId:  186,
					Date:        today,
					YearMonth:   yearMonth,
					ChildrenIds: []int{1860},
					Lines: []model.Line{
						{
							ProductId:    "TST",
							Units:        2,
							ProductPrice: 10.9,
							ChildId:      1860,
						},
					},
					PaymentType: payment_type.BankDirectDebit,
					Note:        "Note 5",
				},
			},
		},
	}

	mockedCfgService := new(mocks.ConfigService)
	mockedCfgService.On("GetString", "yearMonth").Return(yearMonth)

	mockedDbService := new(mocks.DbService)
	mockedDbService.On("FindCustomer", 185).Return(customer185, nil)
	mockedDbService.On("FindCustomer", 186).Return(customer186, nil)
	mockedDbService.On("FindProduct", "TST").Return(productTST, nil)
	mockedDbService.On("FindProduct", "XXX").Return(productXXX, nil)
	mockedDbService.On("FindProduct", "YYY").Return(productYYY, nil)

	mockedOsService := new(mocks.OsService)
	mockedOsService.On("Now").Return(today)

	sut := billingService{
		cfgService: mockedCfgService,
		osService:  mockedOsService,
		dbService:  mockedDbService,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, _ := sut.consumptionsToInvoices(tt.consumptions)
			assert.Equal(t, tt.wantInvoices, got)
		})
	}
}

func Test_billingService_groupConsumptionsByCustomer(t *testing.T) {
	tests := []struct {
		name         string
		consumptions []model.Consumption
		want         map[string][]model.Consumption
	}{
		{
			name:         "group no rectification consumptions",
			consumptions: noRectificationConsumptions,
			want: map[string][]model.Consumption{
				"185": {
					noRectificationConsumptions[0],
					noRectificationConsumptions[1],
					noRectificationConsumptions[2],
					noRectificationConsumptions[3],
				},
				"186": {
					noRectificationConsumptions[4],
					noRectificationConsumptions[5],
					noRectificationConsumptions[6],
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := billingService{}
			got := b.groupConsumptionsByCustomer(tt.consumptions)
			assert.Equal(t, tt.want, got)
		})
	}
}
