package main

import (
	"fmt"
	"unsafe"
)

func intToSlice(num int) []int {
	var res []int

	for num > 0 {
		digit := num % 10
		num /= 10
		res = append([]int{digit}, res...)
	}

	return res
}

func sliceToInt(digits []int) int {
	var num int

	for _, digit := range digits {
		num *= 10
		num += digit
	}

	return num
}

// change the last value with the i
// pop the last element
func removeNoMatterOrder(nums []int, i int) []int { // no matter about the order
	nums[i] = nums[len(nums)-1]
	return nums[:len(nums)-1]
}

// take the left part to the right part
// return without the lastElement
func removePreserveOrder(nums []int, i int) []int { // preserve the order but can change the underlying array.
	copy(nums[i:], nums[i+1:])
	return nums[:len(nums)-1]
}

// The copy function in Go copies elements from a source slice to a destination slice. It fills the destination slice starting from its first element. The function copies up to the lesser of the lengths of the two slices (i.e., the number of elements copied is min(len(dest), len(src))). The copy function works with indices under the hood, copying elements one-by-one from the source to the destination. Returns the len of the smallest slices

// i++ doesn't return any value, if u will try to fmt.Println(i++) <- compile error, missing ,

func iterate(slice []int) {
	if len(slice) == 0 {
		return
	}

	data := unsafe.Pointer(&slice[0]) // addr of the first element in the slice
	uintptrData := uintptr(data)

	for i := 0; i < len(slice); i++ {
		uintptrElem := (uintptrData + uintptr(i)*unsafe.Sizeof(&slice[0]))

		elem := *(*int)(unsafe.Pointer(uintptrElem))
		fmt.Println(elem)
	}
}
