package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const limitWorkers = 100

func main() {
	wrapperLogger := NewWrapperLogger(NewSTDLogger(), 100)

	workCh := make(chan int)

	var wg sync.WaitGroup

	wg.Add(limitWorkers)
	for i := 0; i < limitWorkers; i++ {
		wg.Done()
		go func() {
			for work := range workCh {
				wrapperLogger.Log(fmt.Sprintf("msg: %d", work))
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		workCh <- i
	}
	close(workCh)

	wg.Wait()
	time.Sleep(time.Second * 2)
}

type WrapperLogger struct {
	stdLogger  STDLogger
	mu         sync.Mutex
	isClosedMu sync.RWMutex
	workCh     chan string
	isClosed   bool
	allDone    chan struct{}
}

func (l *WrapperLogger) Log(msg string) {
	l.isClosedMu.RLock()
	defer l.isClosedMu.RUnlock()

	if !l.isClosed {
		l.workCh <- msg
	}
}

func (l *WrapperLogger) doWork() {
	defer func() {
		l.allDone <- struct{}{}
	}()

	for work := range l.workCh {
		go func() {
			l.mu.Lock()
			defer l.mu.Unlock()
			l.stdLogger.Log(work)
		}()
	}
}

func (l *WrapperLogger) Close() {
	defer close(l.allDone)

	if !l.isClosed {
		l.isClosed = true
		close(l.workCh)
	}

	<-l.allDone
}

func NewWrapperLogger(stdLogger STDLogger, buffSize int) WrapperLogger {
	return WrapperLogger{stdLogger: stdLogger, workCh: make(chan string, buffSize), allDone: make(chan struct{})}
}

// std logger

type STDLogger struct{}

func (l *STDLogger) Log(msg string) {
	time.Sleep(time.Second)
	log.Println(msg)
}

func (l *STDLogger) Close() {
	fmt.Println("closed")
}

func NewSTDLogger() STDLogger {
	return STDLogger{}
}
