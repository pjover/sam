package fakes

import (
	"github.com/pjover/sam/internal/domain/ports"
	"time"
)

type osService struct {
}

func FakeOsService() ports.OsService {
	return osService{}
}

func (o osService) OpenUrlInBrowser(url string) error {
	//TODO implement me
	panic("implement me")
}

func (o osService) RunCommand(command string, args ...string) error {
	//TODO implement me
	panic("implement me")
}

func (o osService) CreateDirectory(dirPath string) error {
	//TODO implement me
	panic("implement me")
}

func (o osService) CopyFile(sourceFilePath string, destinationFilePath string) error {
	//TODO implement me
	panic("implement me")
}

func (o osService) ItemExists(itemPath string) (exists bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (o osService) GetTempDirectory() (dirPath string, err error) {
	//TODO implement me
	panic("implement me")
}

func (o osService) CreateZipFile(zipFilePath string, files []string) error {
	//TODO implement me
	panic("implement me")
}

func (o osService) Now() time.Time {
	return time.Date(2022, 8, 15, 11, 51, 59, 0, time.Local)
}

func (o osService) ListFiles(dir string, ext string) (filenames []string, err error) {
	//TODO implement me
	panic("implement me")
}

func (o osService) ReadFile(filePath string) (content []byte, err error) {
	//TODO implement me
	panic("implement me")
}

func (o osService) WriteFile(dirPath string, filename string, content []byte) (filePath string, err error) {
	//TODO implement me
	panic("implement me")
}
