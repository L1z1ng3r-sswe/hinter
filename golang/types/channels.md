### 1. [Функция merge](#merge-function)  
### 2. [Функция throttle](#throttle-function)  
### 3. [Функция calculateSum](#calculate-sum-function)  
### 4. [Функция fetch](#fetch-function)  
### 5. [Функция channels notification](#channels-notification-function)  
### 6. [Функция unpredictableFunc и predictable](#unpredictable-predictable-functions)  
### 7. [Функция semaphore](#semaphore-function)  
### 8. [Обертка над logger](#multythread-wrapper)
### 9. [Get or Create in map](#get-or-create)
### 10. [Count the number of requests](#num-request)
### 11. [Concurrent write into a slice](#concurrent-use-slice)
### 12. [Zero value](#zero-value)
### 13. [Merge with context](#merge-with-context)
### 14. [Закрытие каналов в Go](#why-close-channels-in-go)

---

### Функция merge <a id="merge-function"></a>

```go
func merge(chans ...<-chan int) <-chan int {
	res := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chans {
		ch := ch
		wg.Add(1)

		go func() {
			defer wg.Done()

			for val := range ch {
				res <- val
			}
		}()
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
func throttle(fn func(), ms time.Duration) func() {
	now := time.Now()

	return func() {
		if time.Now().Sub(now) >= ms { // or time.Since
			now = time.Now()
			fn()
		}
	}
}
```

**Описание**: Функция `throttle` принимает функцию `f` и ограничивает её вызовы через указанный промежуток времени `ms`. Полезно для ограничения частоты вызова функции.

### Функция calculateSum <a id="calculate-sum-function"></a>

```go	
func calculateSum(n int) int {
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := range n+1 {
		i := i

		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- i
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res int = 0
	for num := range ch {
		res += num
	}

	return res
}
```

**Описание**: Функция `calculateSum` запускает несколько горутин, каждая из которых отправляет значения в канал `chDone`. Как только все значения будут отправлены, подсчитывается сумма, и программа завершается.

### Функция fetch <a id="fetch-function"></a>

```go
func fetch(urls []string, limit int) {
	if len(urls) < limit {
		limit = len(urls)
	}

	wg := sync.WaitGroup{}
	chURLs := make(chan string)

	wg.Add(limit)
	for i := 0; i < limit; i++ {
		go func() {
			defer wg.Done()
			for url := range chURLs {
				get(url)
			}
		}()
	}

	for _, url := range urls {
		chURLs <- url
	}
	close(chURLs)

	wg.Wait()
}

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error happened on request making")
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
```

**Описание**: Функция `fetch` выполняет параллельные HTTP-запросы на список URL-адресов с ограничением на количество одновременно обрабатываемых запросов. Она использует каналы для передачи URL-адресов рабочим горутинам и синхронизирует их с помощью `sync.WaitGroup`.

### Функция channels notification <a id="channels-notification-function"></a>

```go
func main() {
	x := make(map[int]int, 1)
	var wg sync.WaitGroup

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	wg.Add(3)

	go func() {
		defer wg.Done()
		<-ch1
		x[1] = 2
		ch2 <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		<-ch2
		x[1] = 5
		ch3 <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		<-ch3
		x[1] = 10
	}()

	ch1 <- struct{}{}

	wg.Wait()

	fmt.Println("x[1] =", x[1]) // Expected to see 10
}
```

**Описание**: Функция `main` демонстрирует использование каналов и `WaitGroup` для синхронизации горутин, изменяющих значение в мапе `x`. Горутин выполнены последовательно, изменяя значение `x[1]` по мере передачи сигналов через каналы.

### Функция unpredictableFunc и predictable <a id="unpredictable-predictable-functions"></a>

```go
func main() {
	rand.NewSource(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	dur := rand.Int63n(5000)
	time.Sleep(time.Duration(dur) * time.Millisecond)
	return dur
}

func predictable(timeout time.Duration) int64 {
	ch := make(chan int64, 1)
	go func() {
		ch <- unpredictableFunc()
		close(ch)
	}()

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	select {
	case dur := <-ch:
		return dur
	case <-ctx.Done():
		return 0
	}
}
```

**Описание**: Функция `unpredictableFunc` возвращает случайную длительность, задерживая выполнение программы на это время. Функция `predictable` использует контекст с тайм-аутом для предсказуемого выполнения, возвращая либо результат `unpredictableFunc`, либо 0, если произошел тайм-аут.

### Функция semaphore <a id="semaphore-function"></a>

```go

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	maxWorkers := 10
	ch := make(chan struct{}, maxWorkers)

	for i := 0; i < 100; i++ {
		i := i
		ch <- struct{}{}
		go worker(i, ch, r)
	}
}

func worker(id int, ch chan struct{}, r *rand.Rand) {
	sleepDur := r.Int63n(4)
	time.Sleep(time.Second * time.Duration(sleepDur))
	fmt.Println(id, "done, time: ", sleepDur)
	<-ch
}
```

**Описание**: Функция `worker` использует семафор для ограничения количества одновременно выполняемых горутин. В `main` функции запускается 5 горутин, но только 3 из них могут выполняться одновременно, благодаря использованию буферизированного канала `sem`.

### Многопоточная обертка над logger <a id="multythread-wrapper"></a>

```go

type WrappedLogger struct {
	logger   *FileLogger
	msgBuff  chan string
	closeMu  sync.RWMutex
	done     chan struct{}
	isClosed bool
}

func NewWrappedLogger(logger *FileLogger, buffSize int) WrappedLogger {
	wrappedLogger := WrappedLogger{
		logger:  logger,
		msgBuff: make(chan string, buffSize),
		done:    make(chan struct{}),
	}

	go func() {
		defer close(wrappedLogger.done)

		for msg := range wrappedLogger.msgBuff {
			logger.Log(msg)
		}
	}()

	return wrappedLogger
}

func (l *WrappedLogger) Log(msg string) error {
	l.closeMu.RLock()
	defer l.closeMu.RUnlock()

	if !l.isClosed {
		l.msgBuff <- msg
	}

	return nil
}

func (l *WrappedLogger) Close() {
	l.closeMu.Lock()
	defer l.closeMu.Unlock()

	if !l.isClosed {
		l.isClosed = true
		close(l.msgBuff)
		<-l.done
		l.logger.Close()
	}
}
```

**Описание**: Определяем структуру `WrappedLogger`, которая оборачивает `FileLogger` для обеспечения буферизированного и потокобезопасного логирования. Он обрабатывает сообщения асинхронно с использованием буферизированного канала (`msgBuff`) и гарантирует безопасное закрытие с помощью `sync.RWMutex`. Метод `Close` аккуратно завершает работу логгера, очищая буфер и освобождая ресурсы.

### Get or create in a map <a id="get-or-create"></a>

```go
type SafeMap struct {
	data map[string]string
	rwm sync.RWMutex
}

func (s *SafeMap) Read(key string) string, bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.data[key]
}

func (s *SafeMap) Write(key string, value string) string {
	s.rwm.Lock()	
	defer s.rwm.Unlock()

	val, ok := s.data[key]
	if ok {
		return val
	}

	s.data[key] = value
	return value
}

func (s *SageMap) GetOrCreate(key, value string) string {
	if val, ok := s.Read(key); ok {
		return val
	}

	return s.Write(key, value)
}
```

### Count the number of requests <a id="num-request"></a>

```go
var count int64

const numRequests = 1000

func main() {
	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			doWork()
		}()
	}

	wg.Wait()
	fmt.Println(atomic.LoadInt64(&count))
}

func doWork() {
	time.Sleep(time.Millisecond * 200) // Simulate a network request
	atomic.AddInt64(&count, 1)
}
```

### Concurrent write into a slice <a id="concurrent-use-slice"></a>

```go
func main() {
	n := 200
	c := []int{} // cap = 0, len = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _ = range n {
		wg.Add(1)

		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()

			c = append(c, 1)
		}()
	}

	wg.Wait()

	fmt.Println(len(c))
}
```

### Zero Value <a id="zero-value"></a>

```go
func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("done")
		ch <- 999
	}()

	go func() {
		time.Sleep(time.Second)
		close(ch)

		val, ok := <-ch
		fmt.Printf("Value: %d, Channel Open: %t\n", val, ok)
	}()

	time.Sleep(2 * time.Second)
}

// the output here is:
// done
// Value: 0, Channel Open: false
// panic: send on closed channel

// goroutine 18 [running]:
// main.main.func1()
//         /mnt/c/Users/krasa/OneDrive/Desktop/backup/interviews/main.go:13 +0x65
// created by main.main in goroutine 1
//         /mnt/c/Users/krasa/OneDrive/Desktop/backup/interviews/main.go:11 +0x66
// exit status 2
```

###  <a id="merge-with-context"></a>

```go
func merge(ctx context.Context, chs ...<-chan int) <-chan int {
	var res = make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)
		ch := ch

		go func() {
			defer wg.Done()

			for {
				select  {
					case num, ok := <-ch:
						if !ok {
							return
						}
						res <-num
					case <-ctx.Done():
						return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
```


### Закрытие каналов в Go <a id="why-close-channels-in-go"></a>

**Причина закрытия каналов**:
1. **Избежание утечек горутин**: Закрытие каналов позволяет завершить горутины, которые ожидают данные из канала. Если канал не закрыт, горутины могут оставаться в блокированном состоянии.
2. **Оповещение через `range`**: При использовании `range` для чтения из канала цикл завершится, только если канал закрыт.
3. **Использование `select`**: Закрытие каналов помогает корректно завершать ожидание данных через конструкцию `select`.

ch = make(chan int, 0) - is possible operation