package main

import "fmt"

type Worker interface {
	Do()
}

type User struct{}

func (u *User) Do() { // if not specify the output is: true true false panic
	fmt.Println("Doing...")
}

func main() {
	var w Worker
	fmt.Println(w == nil) // true

	var u *User
	fmt.Println(u == nil) // true

	w = u
	fmt.Println(w == nil) // false

	w.Do() // Doing...

	// var u2 User            // no pointer
	// fmt.Println(u2 == nil) // compile error
}
