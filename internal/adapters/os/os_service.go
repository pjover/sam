package os

import (
	"archive/zip"
	"errors"
	"github.com/pjover/sam/internal/domain/ports"
	"io"
	"os"
	"os/exec"
	"path"
	"runtime"
	"time"
)

type osService struct {
}

func NewOsService() ports.OsService {
	return osService{}
}

func (o osService) OpenUrlInBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return errors.New("unsupported platform")
	}
}

func (o osService) RunCommand(command string, args ...string) error {
	_, err := exec.Command(command, args...).Output()
	return err
}

func (o osService) CreateDirectory(dirPath string) error {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (o osService) ItemExists(itemPath string) (bool, error) {
	_, err := os.Stat(itemPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (o osService) GetTempDirectory() (string, error) {
	dirPath := path.Join(os.TempDir(), "sam")
	exists, err := o.ItemExists(dirPath)
	if err != nil {
		return "", err
	}
	if exists {
		return dirPath, nil
	}
	if err := o.CreateDirectory(dirPath); err != nil {
		return "", err
	}
	return dirPath, nil
}

// CreateZipFile compresses one or many files into a single zip archive file.
// Param 1: zipFilePath is the output zip file's absolute path.
// Param 2: files is a list of files absolute path to add to the zip.
// From: https://golangcode.com/create-zip-files-in-go/
func (o osService) CreateZipFile(zipFilePath string, files []string) error {

	newZipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer func(newZipFile *os.File) {
		err := newZipFile.Close()
		if err != nil {
			panic(err)
		}
	}(newZipFile)

	zipWriter := zip.NewWriter(newZipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			panic(err)
		}
	}(zipWriter)

	// Add files to zip
	for _, file := range files {
		if err = o.addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func (o osService) addFileToZip(zipWriter *zip.Writer, filePath string) error {

	fileToZip, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func(fileToZip *os.File) {
		err := fileToZip.Close()
		if err != nil {
			panic(err)
		}
	}(fileToZip)

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

func (o osService) Now() time.Time {
	return time.Now()
}
