//go:build !windows

package deprecated

import (
	"os/exec"
	"syscall"
)

func _51be323dc33a(_6e6b7ddcf1fb *exec.Cmd) {
	_6e6b7ddcf1fb.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
}
