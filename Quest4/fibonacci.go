package main

import (
	"fmt"
)

var saad = map[int]int{0: 0, 1: 1}

func Fibonacci(a int) int {

	if a < 0 {
		return -1
	}
	if value, exist := saad[a]; exist {
		return value
	}
	saad[a] = Fibonacci(a-1) + Fibonacci(a-2)
	return saad[a]

}

func main() {
	arg1 := 12
	fmt.Println(Fibonacci(arg1))
}
