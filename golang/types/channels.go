// panic: negative counter

func main() {
	wg := &sync.WaitGroup{}

	wg.Done()

	wg.Wait()
}

// undefined behaviour (panic negative or just just leave the program)

func main() {
	wg := &sync.WaitGroup{}

	go operate(*wg)

	wg.Wait()
}

func operate(wg sync.WaitGroup) {
	wg.Done()
}

// deadlock

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go operate(*wg)

	time.Sleep(time.Second)

	wg.Wait()
}

func operate(wg sync.WaitGroup) {
	wg.Done()
}

// will be executed and deadlock at the end

func main() {
	ch := make(chan int, 10)

	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10

	for el := range ch {
		fmt.Println(el)
	}
}

// 10 and deadlock

func main() {
	ch := make(chan int)

	go func() { ch <- 10 }()

	for el := range ch {
		fmt.Println(el)
	}
}

// compiling error because of the i in loop

func main() {
	ch := make(chan int)

	go func() { ch <- 10 }()

	for i, el := range ch {
		fmt.Println(i, el) // error, only el
	}
}

// ok

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	go func() {
		for el := range ch {
			fmt.Println(el)
		}
	}()

	time.Sleep(time.Second)
}

// ok

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	el := <-ch
	fmt.Println(el)

	time.Sleep(time.Second)
}

// ok

func main() {
	ch := make(chan int, 1)

	// go func() {
	ch <- 10
	// }()

	el := <-ch
	fmt.Println(el)

	time.Sleep(time.Second)
}

// deadlock

func main() {
	ch := make(chan int)

	// go func() {
	ch <- 10
	// }()

	el := <-ch
	fmt.Println(el)

	time.Sleep(time.Second)
}

// ok, because of the sender goroutine will block until the receiver goroutine is ready.

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second)
		el := <-ch
		fmt.Println(el)
	}()

	go func() {
		ch <- 10
	}()

	time.Sleep(time.Second)
}

// ok

func main() {
	ch := make(chan int)

	go func() {
		count := 0
		for {
			ch <- count
			count++

			if count%5 == 0 {
				fmt.Println("Pausing for 3 seconds...")
				time.Sleep(3 * time.Second)
				break
			}
		}

		close(ch)
	}()

	for val := range ch {
		fmt.Println("Received:", val)
	}
}

// panic deadlock because of the loop that is waiting for the channel to be closed.

func main() {
	ch := make(chan int)

	go func() {
		count := 0
		for {
			ch <- count
			count++

			if count%5 == 0 {
				fmt.Println("Pausing for 3 seconds...")
				time.Sleep(3 * time.Second)
				return
			}
		}
	}()

	for val := range ch {
		fmt.Println("Received:", val)
	}
}
