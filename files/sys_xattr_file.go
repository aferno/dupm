package files

import(
	"fmt"
	"syscall"
)

func SetXattr() {
	r1, r2, err := syscall.Syscall(syscall.SYS_GETXATTR, 0, 0 ,0)
	fmt.Println(r1, r2, err)
}