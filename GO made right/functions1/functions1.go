// this package is called main because it will be the main entry point of the app
package main

import "fmt"

func Add(a int, b int) int {
	sum := a + b

	return sum
}

func AddInferredReturnType(a int, b int) (sum int) {
	sum = a + b

	return
}

func AddPositives(a, b int) (int, bool) {
	if a < 0 || b < 0 {
		return 0, false
	}

	return a + b, true
}

func main() {
	fmt.Println("Sum of 1 and 2 is", Add(1, 2))

	// check values of function that returns 2 values
	sum, ok := AddPositives(1, 2)
	fmt.Println("Sum of of positives is", sum, "ok", ok)

	sum, ok = AddPositives(-1, 2)
	fmt.Println("Sum of of positives is", sum, "ok", ok)

	x := 99
	sum, ok = 3+7, x > 100
	fmt.Println("sum", sum, "ok", ok)
}
