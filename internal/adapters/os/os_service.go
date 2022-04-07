package os

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/domain/ports"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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

func (o osService) CopyFile(sourceFilePath string, destinationFilePath string) error {
	sourceFileStat, err := os.Stat(sourceFilePath)
	if err != nil {
		return fmt.Errorf("cannot find source file '%s': %s", sourceFilePath, err)
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("source file '%s' is not a regular file", sourceFilePath)
	}
	source, err := os.Open(sourceFilePath)
	if err != nil {
		return fmt.Errorf("cannot open source file '%s': %s", sourceFilePath, err)
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			log.Printf("error closing source file '%s': %s", sourceFilePath, err)
		}
	}(source)

	destination, err := os.Create(destinationFilePath)
	if err != nil {
		return fmt.Errorf("cannot create destination file '%s': %s", destinationFilePath, err)
	}

	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {
			log.Printf("error closing destination file '%s': %s", destinationFilePath, err)
		}
	}(destination)
	_, err = io.Copy(destination, source)
	return err
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

	// GetString the file information
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

func (o osService) ListFiles(dir string, ext string) (filenames []string, err error) {
	err = filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) == ext {
			filenames = append(filenames, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filenames, nil
}

func (o osService) ReadFile(filePath string) (content []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("no s'ha pogut llegir el fitxer '%s': %s", filePath, err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	content, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error llegint el contingut del fitxer '%s': %s", filePath, err)
	}
	return content, nil
}

func (o osService) WriteFile(dirPath string, filename string, content []byte) (filePath string, err error) {
	err = o.CreateDirectory(dirPath)
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut crear el directori %s: %s", dirPath, err)
	}
	filePath = path.Join(dirPath, filename)
	return filePath, os.WriteFile(filePath, content, 0660)
}
