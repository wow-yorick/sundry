package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Status int `json:"status"`
}

func main() {
	var data = []byte(`{"status":200}`)
	result := &Result{}
	if err := json.Unmarshal(data, result); err != nil {
		println("error:", err)
		return
	}

	fmt.Println("result=%+v", result)
}
