package main

import (
	"fmt"
)

func IterativePower(nb int, power int) int {
	f := 1
	for a := 1; a <= power; a++ {
		f *= nb
	}
	return f
}
func main() {
	fmt.Println(IterativePower(4, 3)) //4 its the nb and 3 is the power
}
