package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	fmt.Printf("tick %v", <-tick)
	go func() {
		str, _ := os.Stdin.Read(make([]byte, 1))
		fmt.Println(str)
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			fmt.Println("haha")
		case <-abort:
			fmt.Println("launch aborted!")
			return
		}
	}
	launch()

}

func launch() {
	fmt.Println("go...")
}
