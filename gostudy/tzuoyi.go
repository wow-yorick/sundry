package main

import "fmt"

func main() {

	var tt int64 = 1 << 23
	fmt.Printf("%b", tt)
	ss := []int{1, 2, 3, 4, 5}
	fmt.Println(ss[:0])
}
