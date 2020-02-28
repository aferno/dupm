package files

import(
	"os"
	"fmt"
	"golang.org/x/sys/unix"
)

const DEST_SIZE = 100

func GetXattr(f *os.File) {
	dest := make([]byte, DEST_SIZE)

	fileFileStat, err := f.Stat()
	if err != nil {
		fmt.Printf("There is an error %v\n", err)
	}

	attr, err := unix.Fgetxattr(int(f.Fd()), "user.name", dest)
	if err != nil {
		fmt.Printf("There is nothing in attr %v in %s file\n", err, fileFileStat.Name())
	} else {
		fmt.Printf("Found data - %s in file %s with %dB\n", string(dest), fileFileStat.Name(), attr)
	}
}