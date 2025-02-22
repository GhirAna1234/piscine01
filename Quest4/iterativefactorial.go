package main

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	result := 1
	for a := 1; a <= nb; a++ {
		result *= a
	}
	return result
}
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
