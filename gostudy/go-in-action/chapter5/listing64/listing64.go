package main

import (
	"fmt"
	"sundry/gostudy/go-in-action/chapter5/listing64/counters"
	"sundry/gostudy/go-in-action/chapter5/listing64/entities"
)

func main() {
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)

	// u := entities.User{
	// 	Name: "Bill",
	// 	// email: "bill@email.com",
	// }

	a := entities.Admin{
		Rights: 10,
		// Name:   "xx",
		// Email:  "x@email.com",
	}
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)

}
