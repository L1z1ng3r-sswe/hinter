## Examples

1. [Slice Append Behavior](#slice-append-behavior)
2. [Custom Error Handling](#custom-error-handling)
3. [Zip Two Slices Together](#zip-two-slices-together)
4. [Map Concurrent Access with Mutex](#map-concurrent-access-with-mutex)

---

### Slice Append Behavior <a id="slice-append-behavior"></a>

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

### Custom Error Handling <a id="custom-error-handling"></a>

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

**Explanation:**
- The `customError` struct implements the `Error` interface by defining the `Error` method, which allows us to return it as an `error`.

---

### Zip Two Slices Together <a id="zip-two-slices-together"></a>

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

**Explanation:**
- The `zip` function pairs elements from two slices and returns a slice of pairs.
- It uses the `min` function to ensure that the result is as long as the shorter of the two input slices.

---

### Map Concurrent Access with Mutex <a id="map-concurrent-access-with-mutex"></a>

```go
func main() {
	x := make(map[int]int)
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