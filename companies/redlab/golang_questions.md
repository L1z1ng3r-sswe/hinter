```go
const n = 6

func calculateSum() {
	var res int
	chDone := make(chan int)
	chWait := make(chan struct{})
	defer func() {
		close(chWait)
		close(chDone)
	}()

	for i := 1; i <= n; i++ {
		go func(i int) {
			chDone <- i
		}(i)
	}

	go func() {
		var counter int
		for elem := range chDone {
			res += elem
			counter++
			if counter == n {
				chWait <- struct{}{}
			}
		}
	}()

	<-chWait

	fmt.Println(res)
}

// вам нужно закрыть каналы: чтобы избежать утечек горутины (<-ch), уведомите range, select
```