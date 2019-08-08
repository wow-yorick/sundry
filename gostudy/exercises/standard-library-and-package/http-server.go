package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		//w.Write([]byte("hello"))
		fmt.Fprintf(w, "hello word")
	})
	http.ListenAndServe(":8000", mux)
}
