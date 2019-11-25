package channels

import (
	"fmt"
	"time"
)

func sleep() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx
		}(i)
	}

	fmt.Println(<-ch)           // 输出第一个发送的值
	time.Sleep(2 * time.Second) // 模拟做其他的操作
}
