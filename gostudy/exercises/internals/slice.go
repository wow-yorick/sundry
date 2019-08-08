package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

const M = 10
const N = 5

func printOriginalSlice(subslice *[]int) {
	// data := (*[M]int)(unsafe.Pointer(((*reflect.SliceHeader)(unsafe.Pointer(subslice))).Data))
	data := (*[M]int)(unsafe.Pointer(((*reflect.SliceHeader)(unsafe.Pointer(subslice))).Data))
	fmt.Printf("original\t%p:%+v\n", data, *data)
}

func main() {
	slice := make([]int, M)
	for i, _ := range slice {
		slice[i] = i
	}

	subslice := slice[0:N]

	fmt.Printf("slice\t%p:%+v\n", &slice, slice)
	fmt.Printf("subslice\t%p:%+v\n", &subslice, subslice)
	printOriginalSlice(&subslice)
}
