package files

import (
	"os"
	"fmt"
)

func GetFile(filename string) string {
	filepath := fmt.Sprintf("/dev/shm/%s", filename)
	file_contents, err := os.ReadFile(filepath)
	var file_length = len(file_contents)

	if err != nil{
		panic(fmt.Sprintf("[-] Can't find %s \n", filepath))
	}

	if file_length == 0 {
		panic(fmt.Sprintf("[-] The file is empty, please add some content to it \n"))
	}

	return string(file_contents)
}
