package admin

import (
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/lang"
	"log"
	"path"
	"time"
)

type adminService struct {
	configService ports.ConfigService
	osService     ports.OsService
	langService   lang.LangService
}

func NewAdminService(configService ports.ConfigService, osService ports.OsService, langService lang.LangService) ports.AdminService {
	service := adminService{
		configService: configService,
		osService:     osService,
		langService:   langService,
	}
	_, err := service.CreateWorkingDirectory()
	if err != nil {
		log.Fatal(err)
	}
	return service
}

func (a adminService) Backup() (string, error) {
	fmt.Println("Fent la còpia de seguretat de la base de dades ...")

	tmpDirPath, err := a.osService.GetTempDirectory()
	if err != nil {
		return "", err
	}

	var strSlice = []string{"consumption", "customer", "invoice", "product", "sequence"}
	var files []string
	for _, value := range strSlice {
		fileName := fmt.Sprintf("%s.json", value)
		filePath := path.Join(tmpDirPath, fileName)
		err := a.osService.RunCommand(
			"mongoexport",
			"--db=hobbit",
			fmt.Sprintf("--collection=%s", value),
			fmt.Sprintf("--out=%s", filePath),
		)
		files = append(files, filePath)
		if err != nil {
			return "", err
		}
	}

	zipFilePath, err := a.getZipFilePath()
	if err != nil {
		return "", err
	}

	err = a.osService.CreateZipFile(zipFilePath, files)
	if err != nil {
		return "", err
	}

	return fmt.Sprint("Completada la còpia de seguretat de la base de dades a ", zipFilePath, " ..."), nil
}

func (a adminService) getZipFilePath() (string, error) {
	backupDir := a.configService.GetString("dirs.backup")
	exists, err := a.osService.ItemExists(backupDir)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("el directori %s no existeix", backupDir)
	}
	dateStr := a.osService.Now().Format("060102")
	backupFileName := fmt.Sprintf("%s-Backup.zip", dateStr)
	backupFilePath := path.Join(backupDir, backupFileName)
	return backupFilePath, nil
}

func (a adminService) CreateWorkingDirectory() (string, error) {
	workingTime := a.getWorkingTime()
	yearMonth := model.TimeToYearMonth(workingTime)
	dirName := a.langService.WorkingDir(workingTime)

	dirPath := path.Join(a.configService.GetString("dirs.home"), dirName)
	exists, err := a.osService.ItemExists(dirPath)
	if err != nil {
		return "", err
	}
	if exists {
		_ = a.updateConfig(yearMonth, dirName)
		return a.getInfoText(dirPath), nil
	}

	if err := a.osService.CreateDirectory(dirPath); err != nil {
		return "", err
	}

	if err := a.updateConfig(yearMonth, dirName); err != nil {
		return "", err
	}

	if err := a.osService.OpenUrlInBrowser(dirPath); err != nil {
		return "", err
	}

	return fmt.Sprint("Creat el directori ", dirPath), nil
}

func (a adminService) getInfoText(dirPath string) string {
	days := a.numberOfDaysUntilEndOfMonth()
	var color string
	if days < 2 {
		color = domain.ColorRed
	} else if days < 5 {
		color = domain.ColorYellow
	} else {
		color = domain.ColorGreen
	}

	return fmt.Sprintf("%sSam v%s    %s%s", color, domain.Version, dirPath, domain.ColorReset)
}

func (a adminService) numberOfDaysUntilEndOfMonth() int {
	now := a.osService.Now()
	endOfMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	return int(endOfMonth.Sub(now).Hours() / 24)
}

func (a adminService) getWorkingTime() time.Time {
	var t = a.osService.Now()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

func (a adminService) updateConfig(yearMonth model.YearMonth, dirName string) error {
	if err := a.configService.SetCurrentYearMonth(yearMonth); err != nil {
		return err
	}
	if err := a.configService.SetString("dirs.current", dirName); err != nil {
		return err
	}
	return nil
}
