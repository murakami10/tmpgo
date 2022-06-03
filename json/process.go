package json

import (
	"encoding/json"
	"fmt"
	"log"
)

type GoStruct struct {
	A int    `json:"first"`
	B string `json:"second"`
}

func Process() {
	//fromJson()
	toJson()
}

func fromJson() {
	jsonString := `{"first":1,"second":"bbb"}`

	decode(jsonString)
}

func toJson() {
	stcData := GoStruct{
		A: 1,
		B: "bbb",
	}

	encode(stcData)
}

// json->Go構造体
func decode(jsonString string) {
	var stcData GoStruct

	if err := json.Unmarshal([]byte(jsonString), &stcData); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", stcData)
}

// Go構造体->json
func encode(gs GoStruct) {

	jsonData, err := json.Marshal(gs)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s \n", jsonData)
	return
}
