package main

import (
	"log"
	"tmp/error"
)

func main() {
	u := error.UserService{}
	if err := u.Authenticate(); err != nil {
		log.Fatalf("%+v", err)
	}
}
