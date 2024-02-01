package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	p := Person{"Bapan", 12}
	fmt.Println(p)
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	person_data := json.RawMessage(string(data))
	byte_data, err := person_data.MarshalJSON()
	person := Person{}
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(byte_data, &person)
	fmt.Println(person)
}
