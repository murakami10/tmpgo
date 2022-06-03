package file

import (
	"fmt"
	"os"
)

func File() {
	//open()
	seek()
	//write()
}

func open() {
	fmt.Println("Open method")
	f, err := os.Open("file/text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(count)
	fmt.Println(data, string(data[:count]))

}

func write() {
	fmt.Println("Create method")

	f, err := os.Create("file/write.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	str := "Write this file by Golang!"
	data := []byte(str)
	count, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
}

func seek() {
	f, err := os.Open("file/text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	data := make([]byte, 1024)
	f.Seek(11, 1)
	f.ReadAt(data, 3)
	fmt.Println(string(data))
	fmt.Println()

	f.Seek(0, 1)
	//data, err := ioutil.ReadAll(f)
	f.Read(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
	//
	//f.Seek(11, 1)
	//
	//count, err = f.ReadAt(data, 11)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(count, string(data))
}
