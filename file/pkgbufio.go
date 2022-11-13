package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFileWithBuf() {

	f, err := os.Open("file/text.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	s := bufio.NewScanner(f)
	fmt.Println(s.Text())
	s.Scan()
	fmt.Println(s.Text())

}

func ReadWord() {
	f, err := os.Open("file/text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		fmt.Println(s.Text())
	}
}
