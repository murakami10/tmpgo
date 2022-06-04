package file

import (
	"fmt"
	"net"
)

//func main() {
//	go Listen()
//	time.Sleep(5 * time.Second)
//	go Dial()
//	time.Sleep(5 * time.Second)
//}

func Listen() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("cannot listen", err)
		return
	}

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("cannot accept", err)
	}

	str := "Hello, net pkg"
	data := []byte(str)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("cannot write", err)
	}
}

func Dial() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
	}

	data := make([]byte, 1024)
	count, _ := conn.Read(data)
	fmt.Println(data[:count])
}
