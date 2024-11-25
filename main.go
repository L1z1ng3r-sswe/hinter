package main

import "fmt"

func main() {
	x := 2
	inc(x)

	fmt.Println(x)
}

func inc(i interface{}) int {
	return i.(int)
}