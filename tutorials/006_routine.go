package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i  int64
	mu sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		if i%100 == 0 {
			fmt.Println(i)
		}
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()
	fmt.Println(c.i)
}
