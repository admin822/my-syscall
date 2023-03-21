package mysyscall

import (
	"fmt"
	"os/exec"
	"syscall"
)

func MySetGid(id int) {
	syscall.Setgid(id)
}

func RunScript(cmdStr string) {
	out, err := exec.Command(cmdStr).CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", out)
}
