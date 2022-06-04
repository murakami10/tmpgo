package file

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//func main() {
//	TranslateFile()
//}

func TranslateFile() {
	f, err := os.Open("file/text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	TranslateIntoGerman(f)

}

func TranslateIntoGerman(i io.Reader) {
	data := make([]byte, 1024)
	count, _ := i.Read(data)

	result := strings.ReplaceAll(string(data[:count]), "Hello", "Guten Tag")

	fmt.Println(result)
}
