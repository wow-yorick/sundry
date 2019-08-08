package main

import "fmt"

func main() {
	// i := 0
	// for { //since there are no checks, this is an infinite loop
	// 	if i >= 3 {
	// 		break
	// 	}
	// 	//break out of this for loop when this condition is met
	// 	fmt.Println("Value of i is:", i)
	// 	i++
	// }
	// fmt.Println("A statement just after for loop.")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}

	// LABEL1:
	// 	for i := 0; i <= 5; i++ {
	// 		for j := 0; j <= 5; j++ {
	// 			if j == 4 {
	// 				break LABEL1
	// 			}
	// 			fmt.Printf("i is: %d, and j is: %d\n", i, j)
	// 		}
	// 	}
}
