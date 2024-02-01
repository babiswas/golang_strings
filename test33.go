package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{"name":"Bapan","id":200,"marks":22.2}`
	var obj any
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		fmt.Println(err)
	}
	m := obj.(map[string]any)
	fmt.Println("Displaying person name:")
	fmt.Println(m["name"])
	fmt.Println("Displaying id of the person:")
	fmt.Println(m["id"])
	fmt.Println("Displaying marks of the person:")
	fmt.Println(m["marks"])
}
