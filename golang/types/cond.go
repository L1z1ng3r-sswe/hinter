package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu       sync.Mutex
		cond     = sync.NewCond(&mu)
		queue    []int
		capacity = 5
	)

	// producer
	go func() {
		for _ = range 10 {
			mu.Lock()
			defer mu.Unlock()

			for len(queue) >= capacity {
				cond.Wait()
			}

			queue = append(queue, 9)
			cond.Signal()
		}
	}()

	// consumer
	go func() {
		for {
			mu.Lock()
			defer mu.Unlock()

			for len(queue) <= 0 {
				cond.Wait()
			}

			fmt.Println(queue[0])
			queue = queue[:1]
			cond.Signal()
		}
	}()
}
