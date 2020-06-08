package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	type TestData struct {
		Data []int `json:"data"`
	}
	var entity TestData
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	err = json.Unmarshal(data, &entity)
	if err != nil {
		fmt.Println("Unmarshal error", err)
		return
	}
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	var result = make(map[int]bool)

	for s := range entity.Data {
		result[s] = true
	}

	fmt.Println("Result Map:", result)
}
