1. [Функция intToSlice](#int-to-slice-function)  
2. [Функция sliceToInt](#slice-to-int-function)  
3. [Функция removeNoMatterOrder](#remove-no-matter-order-function)  
4. [Функция removePreserveOrder](#remove-preserve-order-function)  
5. [Функция iterate](#iterate-function)  
6. [Функция customAppend](#custom-append-function)  
7. [Функция isPalindrome](#is-palindrome-function)  
8. [Функция reverse](#reverse-function)
9. [Method copy](#copy-method)

---

### Функция intToSlice <a id="int-to-slice-function"></a>

```go
func intToSlice(num int) []int { // 1234 -> [1,2,3,4]
	var res []int

	for num > 0 {
		lastDigit := num % 10 // pick the curr digit
		res = append([]int{lastDigit}, res...)
		num /= 10 // shift to the next digit
	}

	return res
}
```

**Описание**: Преобразует целое число в срез его цифр.

---

### Функция sliceToInt <a id="slice-to-int-function"></a>

```go
func sliceToInt(digits []int) int { // 1,2,3,4
	var res int

	for _, digit := range digits {
		res *= 10
		res += digit
	}

	return res
}
```

**Описание**: Преобразует срез цифр в целое число.

---

### Функция removeNoMatterOrder <a id="remove-no-matter-order-function"></a>

```go
func removeNoMatterOrder(nums []int, i int) []int { // no matter about the order
	nums[i] = nums[len(nums)-1]
	return nums[:len(nums)-1]
}
```

**Описание**: Удаляет элемент из среза, заменяя его последним элементом, порядок элементов не сохраняется.

---

### Функция removePreserveOrder <a id="remove-preserve-order-function"></a>

```go
func removePreserveOrder(nums []int, i int) []int { // preserve the order but can change the underlying array.
	copy(nums[i:], nums[i+1:])
	return nums[:len(nums)-1]
}
```

**Описание**: Удаляет элемент из среза, сохраняя порядок остальных элементов, но может изменять исходный срез.

---

### Функция iterate <a id="iterate-function"></a>

```go
func iterate(slice []int) {
	if len(slice) == 0 {
		return
	}

	firstAddr := uintptr(unsafe.Pointer(&slice[0])) //  address of the first element in the slice

	for i := 0; i < len(slice); i++ {
		uintptrElem := (firstAddr + uintptr(i)*unsafe.Sizeof(&slice[0]))

		elem := *(*int)(unsafe.Pointer(uintptrElem))
		fmt.Println(elem)
	}
}
```

**Описание**: Итерация по срезу с использованием небезопасных указателей.

---

### Функция customAppend <a id="custom-append-function"></a>

```go
func customAppend[T any](slice []T, newVals ...T) []T {
	newSlice := slice
	newLength := len(slice) + len(newVals)

	if newLength > cap(slice) {
		newCap := newLength

		if newCap < cap(slice)*2 {
			newCap = cap(slice) * 2
		}

		newSlice = make([]T, len(slice), newCap)

		copy(newSlice, slice)
	}

	newSlice = newSlice[:newLength]
	copy(newSlice[len(slice):], newVals)

	return newSlice
}
```

**Описание**: Реализация кастомного `append`, который динамически увеличивает ёмкость среза и добавляет новые элементы.

---

### Функция isPalindrome <a id="is-palindrome-function"></a>

```go
func isPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}
```

**Описание**: Проверяет, является ли строка палиндромом.

---

### Функция reverse <a id="reverse-function"></a>

```go
func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
}
```

**Описание**: Разворачивает срез, меняя порядок элементов на обратный.

---

### Method copy <a id="copy-method"></a>


// The copy function in Go copies elements from a source slice to a destination slice. It fills the destination slice starting from its first element. The function copies up to the lesser of the lengths of the two slices (i.e., the number of elements copied is min(len(dest), len(src))). The copy function works with indices under the hood, copying elements one-by-one from the source to the destination. Returns the len of the smallest slices

```go
func main() {
	sl := make([]int, 2, 5) // [0, 1] 1, 2, 3]
	add(sl)
	fmt.Println(sl) // 0,1
	sl = sl[:5]     // this way works 0,1,1,2,3
}

func add(sl []int) {
	sl[1] = 1
	sl = append(sl, []int{1, 2, 3}...)
}

```