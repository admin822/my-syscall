package mysyscall

import (
	"os/exec"
	"syscall"
)

// wrapper function for syscall's wrapper function
func MySetGid(id int) {
	syscall.Setgid(id)
}

// Wrapper function for os/exec's shell command execution function
func RunScript(cmdStr string) ([]byte, error) {
	out, err := exec.Command(cmdStr).CombinedOutput()
	return out, err
}
