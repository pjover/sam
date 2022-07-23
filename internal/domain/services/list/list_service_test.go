package list

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	modelMocks "github.com/pjover/sam/internal/domain/ports/mocks"
	"github.com/pjover/sam/internal/domain/services/loader"
	loaderMocks "github.com/pjover/sam/internal/domain/services/loader/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_listService_ListChildConsumptions_one_child(t *testing.T) {
	var mockedConfigService = new(modelMocks.ConfigService)

	var mockedDbService = new(modelMocks.DbService)
	mockedDbService.On("FindActiveChildConsumptions", 1480).Return(
		[]model.Consumption{
			model.NewConsumption(
				"1",
				1480,
				"TST",
				2,
				model.NewYearMonth(2022, 7),
				"Note 1",
				false,
				""),
			model.NewConsumption(
				"2",
				1480,
				"XXX",
				1,
				model.NewYearMonth(2022, 7),
				"Note 2",
				false,
				""),
		},
		nil,
	)

	var mockedBulkLoader = new(loaderMocks.BulkLoader)
	mockedBulkLoader.On("LoadProducts").Return(
		map[string]model.Product{
			"TST": model.ProductTST,
			"XXX": model.ProductXXX,
		},
		nil,
	)

	mockedDbService.On("FindChild", 1480).Return(
		model.TestChild1480,
		nil,
	)

	type fields struct {
		configService ports.ConfigService
		dbService     ports.DbService
		bulkLoader    loader.BulkLoader
	}
	type args struct {
		childId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "One child consumptions",
			fields: fields{
				configService: mockedConfigService,
				dbService:     mockedDbService,
				bulkLoader:    mockedBulkLoader,
			},
			args: args{
				childId: 1480,
			},
			want: "Laura Llull (1480): 30.90 €\n" +
				"  [2022-07]   2.0 x TST :   21.80\n" +
				"  [2022-07]   1.0 x XXX :    9.10\n",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := NewListService(
				tt.fields.configService,
				tt.fields.dbService,
				tt.fields.bulkLoader,
			)
			got, err := sut.ListChildConsumptions(tt.args.childId)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
