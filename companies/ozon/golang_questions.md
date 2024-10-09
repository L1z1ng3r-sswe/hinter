# Example 1: Slice Append Behavior

```go
func main() {
	nums := []int{1, 2, 3}

	addNum(nums[0:2]) 
	fmt.Println(nums) // Output: [1 2 3]

	addNums(nums[0:2]) 
	fmt.Println(nums)  // Output: [1 2 3]
}

func addNum(nums []int) {
	nums = append(nums, 4) 
}

func addNums(nums []int) {
	nums = append(nums, 5, 6) 
}
```

---

# Example 2: Custom Error Handling

```go
func handle() error {
  err := &customError{msg: "some error"}

  return err
}

type customError struct {
  msg string
}

func (ce customError) Error() string {
  return ce.msg
}
```

Explanation:
- The `customError` struct implements the `Error` interface by defining the `Error` method, which allows us to return it as an `error`.

---

# Example 3: Zip Two Slices Together

```go
func zip(s1, s2 []int) [][]int {
	minLen := min(len(s1), len(s2))

	res := make([][]int, 0, minLen)

	for i := 0; i < minLen; i++ {
		res = append(res, []int{s1[i], s2[i]})
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
```

Explanation:
- The `zip` function pairs elements from two slices and returns a slice of pairs.
- It uses the `min` function to ensure that the result is as long as the shorter of the two input slices.

---

# Example 4: Map Concurrent Access with Mutex

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	x := make(map[int]int, 1)
	var mu sync.Mutex
	var wg sync.WaitGroup

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
  
	wg.Add(3)

	go func() {
		defer wg.Done()
		<-ch
		mu.Lock()
		defer mu.Unlock()
		x[1] = 2
		ch <- struct{}{} 
	}()

	go func() {
		defer wg.Done()
		<-ch
		mu.Lock()
		defer mu.Unlock()
		x[1] = 5
		ch <- struct{}{} 
	}()

	go func() {
		defer wg.Done()
		<-ch
		mu.Lock()
		defer mu.Unlock()
		x[1] = 10
		close(ch) 
	}()

	ch <- struct{}{}

	wg.Wait()

	fmt.Println("x[1] =", x[1]) // Expected to see 10
}
```
