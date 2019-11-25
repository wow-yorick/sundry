package channels

import "fmt"

func what() {
	var x int
	x = 1
	ch := make(chan int)
	go func() {
		ch <- x
	}()
	fmt.Println("channel:", <-ch)
	close(ch)
}
