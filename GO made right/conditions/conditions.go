package main

import (
	"fmt"
	"math/rand"
)

func GetNumber() int {
	return rand.Intn(4)
}

func main() {
	x := GetNumber()

	if x := GetNumber(); x > 1 {
		fmt.Println("It is biggish")
	} else {
		fmt.Println("Small")
	}

	fmt.Println("X is", x)
}
