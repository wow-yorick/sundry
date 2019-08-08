package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = []byte("123")
	s := *(*string)(unsafe.Pointer(&b))

	b[1] = '4'
	fmt.Printf("%+v\n", s)
}
