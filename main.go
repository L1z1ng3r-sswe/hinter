package main

import "fmt"

func main() {
	str := "こんにちは"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}
