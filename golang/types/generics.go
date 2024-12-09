package main

import "fmt"

type Stack[T any] struct {
	queue []T
}

type Woman struct {
	name string
}

type Man struct {
	name string
}

type NoBinary struct {
	name string
}

type UnKnown struct {
	name string
}

type Jalap[T Woman | Man | NoBinary | UnKnown] struct {
	price int
}

func (j *Jalap[T]) Fuck() T {
	var zeroVal T
	return zeroVal
}

func (s *Stack[T]) EnQueue(val T) {
	s.queue = append(s.queue, val)
}

func (s *Stack[T]) DeQueue() (T, bool) {
	if len(s.queue) <= 0 {
		var zeroValue T
		return zeroValue, false
	}

	val := s.queue[len(s.queue)-1]
	s.queue = s.queue[:len(s.queue)-1]
	return val, true
}

func main() {
	stack := Stack[int]{}

	stack.EnQueue(1)
	stack.EnQueue(3)
	stack.EnQueue(10)
	stack.EnQueue(11)

	for len(stack.queue) > 0 {
		fmt.Println(stack.DeQueue())
	}
}
