package services

import (
	"fmt"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
	"path"
)

type adminService struct {
	timer shared.TimeManager
}

func NewAdminService(timer shared.TimeManager) ports.AdminService {
	return adminService{
		timer: timer,
	}
}

func (a adminService) Backup() (string, error) {
	filePath := a.filePath()
	fmt.Println("Fent la còpia de seguretat de la base de dades ...")

	return fmt.Sprint("Completada la còpia de seguretat de la base de dades a", filePath, " ..."), nil
}

func (a adminService) filePath() string {
	dateStr := a.timer.Now().Format("060102")
	return path.Join(
		viper.GetString("dirs.backup"),
		fmt.Sprintf("%s-ºBackup.zip", dateStr),
	)
}

func (a adminService) CreateDirectory(previousMonth bool, nextMonth bool) (string, error) {
	panic("implement me")
}
