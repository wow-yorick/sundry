package other

import "fmt"

func testcop() {
	// 	slice1 := []int{1, 2, 3, 4, 5}
	// 	slice2 := []int{6, 7, 8}
	// 	// c2 := copy(slice2, slice1)
	// 	// fmt.Printf("slice1: %v slice2: %v c1:%v \n", slice1, slice2, c2)
	// 	c1 := copy(slice1, slice2)
	// 	fmt.Printf("slice1: %v slice2: %v c1:%v \n", slice1, slice2, c1)
	var s []byte
	if s == nil {
		fmt.Println("s == nil")
	} else {
		fmt.Println("s != nil")
	}
}
