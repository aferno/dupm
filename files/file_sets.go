package files

import(
	"fmt"
	"encoding/hex"
)

var sha1Files = map[string][]string{}

func FileSetBuilder(checkSum []byte, fileName string) {
	encodedStr := hex.EncodeToString(checkSum)
	if _, ok := sha1Files[encodedStr]; ok {
		fmt.Println("File already exists", fileName)		
	}
	sha1Files[encodedStr] = append(sha1Files[encodedStr], fileName)
}

func hardLink() {
	//TBD
}