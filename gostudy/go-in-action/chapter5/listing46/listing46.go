package main

import "fmt"

type duration int

func (d duration) pretty() string {
	return fmt.Sprintf("Duration: %d", d)
}

func main() {
	s := duration(42).pretty()
	fmt.Println(s)
}
