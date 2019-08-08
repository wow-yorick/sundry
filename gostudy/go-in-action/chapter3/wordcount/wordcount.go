package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sundry/gostudy/go-in-action/chapter3/words"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("there was en error opening the file:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
