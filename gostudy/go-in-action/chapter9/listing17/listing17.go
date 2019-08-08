package main

import (
	"log"
	"net/http"
	"sundry/gostudy/go-in-action/chapter9/listing17/handlers"
)

func main() {
	handlers.Routes()
	log.Println("listener : Started : listening on :4000")
	http.ListenAndServe(":4000", nil)
}
