package mysyscall

import (
	"syscall"
)

func MySetGid(id int) {
	syscall.Setgid(id)
}
