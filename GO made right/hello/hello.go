package main

import (
	"fmt"

	"wiltech.com/data"
)

const (
	ONE = 1
	TWO = 2
)
const (
	B_1 = 1 << iota
	B_2
	B_4
	B_8
)

func main() {
	// var message string = "Hello, World!"
	var message string // this will be am empty string as default
	// var message = "Hello, World!"
	fmt.Println(message)
	fmt.Printf("message is: %s, lenght is %d\n", message, len(message))
	fmt.Println("message from other module is ", data.Message)

	var x int = -3
	fmt.Printf("x is %v, %d, %T, at %p\n ", x, x, x, &x)

	fmt.Printf("1, 2, 4, 8... %d, %d, %d, %d\n", B_1, B_2, B_4, B_8)
}
