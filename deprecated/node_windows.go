//go:build windows

package deprecated

import (
	"os/exec"
	"syscall"
)

const (
	_633fafe0cbf6 uint32 = 0x00000200
	_c1956a018915 uint32 = 0x08000000
)

func _51be323dc33a(_01ba18b5dcb4 *exec.Cmd) {
	_01ba18b5dcb4.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
		CreationFlags: _633fafe0cbf6 |
			_c1956a018915,
	}
}
