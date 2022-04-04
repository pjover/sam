package billing

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/sequence_type"
	"github.com/pjover/sam/internal/domain/ports/mocks"
	"github.com/pjover/sam/test/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var yearMonth, _ = model.StringToYearMonth("2022-02")

var today = time.Date(2022, 2, 16, 20, 33, 59, 0, time.Local)

var noRectificationConsumptions = []model.Consumption{
	{
		Id:        "AA1",
		ChildId:   1480,
		ProductId: "TST",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 1",
	},
	{
		Id:        "AA2",
		ChildId:   1480,
		ProductId: "TST",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 2",
	},
	{
		Id:        "AA3",
		ChildId:   1481,
		ProductId: "TST",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 3",
	},
	{
		Id:        "AA4",
		ChildId:   1481,
		ProductId: "XXX",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 4",
	},
	{
		Id:        "AA5",
		ChildId:   1490,
		ProductId: "TST",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "Note 5",
	},
	{
		Id:        "AA7",
		ChildId:   1490,
		ProductId: "YYY",
		Units:     2,
		YearMonth: yearMonth,
		Note:      "",
	},
	{
		Id:        "AA8",
		ChildId:   1490,
		ProductId: "YYY",
		Units:     -2,
		YearMonth: yearMonth,
		Note:      "",
	},
}

var rectificationConsumptions = []model.Consumption{
	{
		Id:              "AA1",
		ChildId:         1480,
		ProductId:       "TST",
		Units:           -2,
		YearMonth:       yearMonth,
		Note:            "Note r1",
		IsRectification: true,
	},
	{
		Id:              "AA2",
		ChildId:         1480,
		ProductId:       "TST",
		Units:           -2,
		YearMonth:       yearMonth,
		Note:            "Note r2",
		IsRectification: true,
	},
}

var sequences = []model.Sequence{
	{
		sequence_type.StandardInvoice,
		188,
	},
	{
		sequence_type.RectificationInvoice,
		10,
	},
}

func Test_BillConsumptions_without_rectification(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr error
	}{
		{
			name: "BillConsumptions",
			want: " 1. Cara Santamaria 148  F-189  2022-02    83.60  Rebut  2.0 TST (21.80),2.0 XXX (18.20),4.0 TST (43.60)\n" +
				"Total 1 Rebut: 83.60 €\n" +
				" 1. Joana Petita 149  X-1  2022-02    21.80  Tranferència  2.0 TST (21.80)\n" +
				"Total 1 Tranferència: 21.80 €\n" +
				"TOTAL: 105.40 €\n",
			wantErr: nil,
		},
	}

	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("GetCurrentYearMonth").Return(yearMonth)

	mockedDbService := new(mocks.DbService)
	mockedDbService.On("FindAllActiveConsumptions").Return(noRectificationConsumptions, nil)
	mockedDbService.On("FindCustomer", 148).Return(test_data.Customer148, nil)
	mockedDbService.On("FindCustomer", 149).Return(test_data.Customer149, nil)
	mockedDbService.On("FindProduct", "TST").Return(test_data.ProductTST, nil)
	mockedDbService.On("FindProduct", "XXX").Return(test_data.ProductXXX, nil)
	mockedDbService.On("FindProduct", "YYY").Return(test_data.ProductYYY, nil)
	mockedDbService.On("FindAllSequences").Return(sequences, nil)
	mockedDbService.On("InsertInvoices", mock.Anything).Return(nil)
	mockedDbService.On("UpdateConsumptions", mock.Anything).Return(nil)
	mockedDbService.On("UpdateSequences", mock.Anything).Return(nil)

	mockedOsService := new(mocks.OsService)
	mockedOsService.On("Now").Return(today)

	sut := billingService{
		configService: mockedConfigService,
		osService:     mockedOsService,
		dbService:     mockedDbService,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sut.BillConsumptions()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_BillConsumptions_with_rectification(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr error
	}{
		{
			name: "BillConsumptions",
			want: " 1. Cara Santamaria 148  F-189  2022-02    83.60  Rebut  2.0 TST (21.80),2.0 XXX (18.20),4.0 TST (43.60)\n" +
				" 2. Cara Santamaria 148  R-11  2022-02   -43.60  Rebut  -4.0 TST (-43.60)\n" +
				"Total 2 Rebut: 40.00 €\n" +
				" 1. Joana Petita 149  X-1  2022-02    21.80  Tranferència  2.0 TST (21.80)\n" +
				"Total 1 Tranferència: 21.80 €\n" +
				"TOTAL: 61.80 €\n",
			wantErr: nil,
		},
	}

	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("GetCurrentYearMonth").Return(yearMonth)

	mockedDbService := new(mocks.DbService)
	mockedDbService.On("FindAllActiveConsumptions").Return(append(noRectificationConsumptions, rectificationConsumptions...), nil)
	mockedDbService.On("FindCustomer", 148).Return(test_data.Customer148, nil)
	mockedDbService.On("FindCustomer", 149).Return(test_data.Customer149, nil)
	mockedDbService.On("FindProduct", "TST").Return(test_data.ProductTST, nil)
	mockedDbService.On("FindProduct", "XXX").Return(test_data.ProductXXX, nil)
	mockedDbService.On("FindProduct", "YYY").Return(test_data.ProductYYY, nil)
	mockedDbService.On("FindAllSequences").Return(sequences, nil)
	mockedDbService.On("InsertInvoices", mock.Anything).Return(nil)
	mockedDbService.On("UpdateConsumptions", mock.Anything).Return(nil)
	mockedDbService.On("UpdateSequences", mock.Anything).Return(nil)

	mockedOsService := new(mocks.OsService)
	mockedOsService.On("Now").Return(today)

	sut := billingService{
		configService: mockedConfigService,
		osService:     mockedOsService,
		dbService:     mockedDbService,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sut.BillConsumptions()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
