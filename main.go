package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	x := make(map[int]int, 1)
	var mu = sync.Mutex{}

	var wg sync.WaitGroup

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	wg.Add(3)

	go func() {
		wg.Done()
		<-ch1

		mu.Lock()
		defer mu.Unlock()
		x[1] = 2

		ch2 <- struct{}{}
		close(ch1)
	}()

	go func() {
		wg.Done()
		<-ch2

		mu.Lock()
		defer mu.Unlock()
		x[1] = 5

		ch3 <- struct{}{}
		close(ch2)

	}()

	go func() {
		wg.Done()
		<-ch3

		mu.Lock()
		defer mu.Unlock()
		x[1] = 10

		close(ch3)
	}()

	wg.Wait()
	ch1 <- struct{}{}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("x[1] =", x[1]) // assume to see 10
}
