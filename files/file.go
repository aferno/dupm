package files

import(
	"fmt"
	"os"
	"path/filepath"
	"crypto/sha1"
)

func WalkInRoot(source string) {
	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			checkSum(path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error during walking %v\n", err)
		return
	}

}

func checkSum(fileName string) {

	f, err := os.Open(fileName)

	defer f.Close()

	fileStat, err := f.Stat(); 
	if err != nil {
		fmt.Printf("Error when process file %v\n", err)
	}
	fileSize := fileStat.Size()
	data := make([]byte, fileSize)
	bytesRead, err := f.Read(data)
	if err != nil {
		fmt.Printf("Error when process file %v %d\n", err, bytesRead)
	}
	if err != nil {
		fmt.Printf("Error when process file %v\n", err)
	}
	sha1 := sha1.Sum([]byte(data))
	FileSetBuilder(sha1[:12], fileName)
	GetXattr(f)

}