package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	type TestData struct {
		Data []int
	}
	var entity TestData
	data, err := ioutil.ReadFile("test.json")
	json.Unmarshal([]byte(data), &entity)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	var result = make(map[string][]int)
	result["data"] = entity.Data
	fmt.Println("Contents of file:", result["data"])
}
