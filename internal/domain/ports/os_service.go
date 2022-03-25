package ports

import "time"

type OsService interface {
	OpenUrlInBrowser(url string) error
	RunCommand(command string, args ...string) error
	CreateDirectory(dirPath string) error
	ItemExists(itemPath string) (bool, error)
	GetTempDirectory() (string, error)
	CreateZipFile(zipFilePath string, files []string) error
	Now() time.Time
	ListFiles(dir string, ext string) (filenames []string, err error)
}
