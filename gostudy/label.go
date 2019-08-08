package main

import (
	"fmt"
)

func main() {
	a := 1
Loop:
	for a < 10 {
		if a == 5 {
			a++
			goto Loop
		}
		fmt.Println(a)
		a++
	}

	fmt.Println(a)
}
