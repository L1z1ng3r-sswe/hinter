### 1. [Функция merge](#merge-function)  
### 2. [Функция throttle](#throttle-function)  
### 3. [Функция calculateSum](#calculate-sum-function)  
### 4. [Закрытие каналов в Go](#why-close-channels-in-go)

---

### Функция merge <a id="merge-function"></a>

```go
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
```

**Описание**: Функция `merge` принимает несколько каналов на чтение (`<-chan int`) и объединяет их в один результирующий канал `res`, возвращая все значения из переданных каналов.

### Функция throttle <a id="throttle-function"></a>

```go
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
```

**Описание**: Функция `throttle` принимает функцию `f` и ограничивает её вызовы через указанный промежуток времени `ms`. Полезно для ограничения частоты вызова функции.

### Функция calculateSum <a id="calculate-sum-function"></a>

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
```

**Описание**: Функция `calculateSum` запускает несколько горутин, каждая из которых отправляет значения в канал `chDone`. Как только все значения будут отправлены, подсчитывается сумма, и программа завершается.

### Закрытие каналов в Go <a id="why-close-channels-in-go"></a>

**Причина закрытия каналов**:
1. **Избежание утечек горутин**: Закрытие каналов позволяет завершить горутины, которые ожидают данные из канала. Если канал не закрыт, горутины могут оставаться в блокированном состоянии.
2. **Оповещение через `range`**: При использовании `range` для чтения из канала цикл завершится, только если канал закрыт.
3. **Использование `select`**: Закрытие каналов помогает корректно завершать ожидание данных через конструкцию `select`.