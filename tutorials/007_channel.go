package main

import (
	"fmt"
	"time"
)

func main() {
    myChannel := make(chan int)

    go func() {
        defer close(myChannel)
        for i := 0; i < 60; i++ {
	    time.Sleep(1*time.Second)
            myChannel <- i
        }
    }()

    for data := range myChannel {
        fmt.Println(data)
    }
}
