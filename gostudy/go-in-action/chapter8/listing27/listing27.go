package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var JSON = `{
"name":"xxx",
"title":"aaa",
"contact" :{
"home":"23",
"cell":"324545"
}
}`

func main() {
	//var c Contact
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(c)
	fmt.Println("Name", c["name"])
	fmt.Println("Title", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}
