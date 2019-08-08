package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "ag"
	c["title"] = "pro"
	c["contact"] = map[string]interface{}{
		"home": "1",
		"cell": "3",
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))
}
