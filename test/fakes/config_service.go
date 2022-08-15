package fakes

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"time"
)

type configService struct {
}

func FakeConfigService() ports.ConfigService {
	return configService{}
}

func (c configService) Init() {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetString(key string) string {
	//TODO implement me
	panic("implement me")
}

func (c configService) SetString(key string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetTime(key string) time.Time {
	//TODO implement me
	panic("implement me")
}

func (c configService) SetTime(key string, value time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetCurrentYearMonth() model.YearMonth {
	return model.NewYearMonth(2022, 8)
}

func (c configService) SetCurrentYearMonth(yearMonth model.YearMonth) error {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetConfigDirectory() string {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetHomeDirectory() string {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetWorkingDirectory() string {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetInvoicesDirectory() string {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetReportsDirectory() string {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetCustomersCardsDirectory() string {
	//TODO implement me
	panic("implement me")
}

func (c configService) GetBackupDirectory() string {
	//TODO implement me
	panic("implement me")
}
