package main

import (
	"log"
	"sundry/gostudy/go-in-action/chapter7/patterns/search"
)

func main() {
	results := search.Submit(
		"golang",
		search.OnlyFirst,
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : results : info : %+v\n", result)
	}

	// results := search.Submit("golang", search.Google, search.Bing, search.Yahoo)

	// for _, result := range results {
	// 	log.Printf("main : results : info : %+v\n", result)
	// }
}
