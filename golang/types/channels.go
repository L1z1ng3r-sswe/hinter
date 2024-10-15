// merge

func merge(chans ...<-chan int) chan int {
	res := make(chan int)
	var wg sync.WaitGroup

	receiver := func(ch <-chan int) {
		defer wg.Done()

		for elem := range ch {
			res <- elem
		}
	}

	for _, ch := range chans {
		wg.Add(1)
		go receiver(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

// throttle

func throttle(f func(), ms time.Duration) func() {
	var lastTime time.Time

	return func() {
		now := time.Now()
		if now.Sub(lastTime) < ms {
			lastTime = now
			f()
		}
	}
}
