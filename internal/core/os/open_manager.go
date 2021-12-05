package os

import (
	"errors"
	"os/exec"
	"runtime"
)

type OpenManager interface {
	OnDefaultApp(url string) error
}

type openManager struct {
}

func NewOpenManager() OpenManager {
	return openManager{}
}

func (o openManager) OnDefaultApp(url string) error {
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
