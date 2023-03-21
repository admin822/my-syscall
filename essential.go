package mysyscall

import (
	"os/exec"
	"syscall"
)

func MySetGid(id int) {
	syscall.Setgid(id)
}

func RunScript(cmdStr string) ([]byte, error) {
	out, err := exec.Command(cmdStr).CombinedOutput()
	return out, err
}
