package main

import (
	"fmt"
	"time"
)

func addChannelnum(ch chan int, num int) {
	ch <- num
}
func main() {
	myChannel := make(chan int)

	go func() {
		defer close(myChannel)

		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			addChannelnum(myChannel, i)
		}
	}()

	for data := range myChannel {
		fmt.Println(data)
	}
}
