package main

import (
	"log"
	"os"
	_ "sundry/gostudy/go-in-action/chapter2/sample/matchers"
	"sundry/gostudy/go-in-action/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
