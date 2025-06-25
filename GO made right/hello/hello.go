package main

import (
	"fmt"

	"wiltech.com/data"
)

func main() {
	// var message string = "Hello, World!"
	var message string // this will be am empty string as default
	// var message = "Hello, World!"
	fmt.Println(message)
	fmt.Printf("message is: %s, lenght is %d\n", message, len(message))
	fmt.Println("message from other module is ", data.Message)
}
