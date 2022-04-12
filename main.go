package main

import (
	"fmt"
	"os"
)

type aaa string

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
	fmt.Println(data, string(data[:count]))

	fmt.Println()
	fmt.Println("Create method")
	fmt.Println()

	f, err = os.Create("write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	str := "Write this file by Golang!"
	data = []byte(str)
	count, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
}
