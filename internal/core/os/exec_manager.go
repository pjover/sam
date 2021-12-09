package os

import (
	"errors"
	"os/exec"
	"runtime"
)

type ExecManager interface {
	BrowseTo(url string) error
	Run(command string, args ...string) ([]byte, error)
}

type execManager struct {
}

func NewExecManager() ExecManager {
	return execManager{}
}

func (o execManager) BrowseTo(url string) error {
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

func (o execManager) Run(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}
