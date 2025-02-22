package main

import (
	"fmt"
)

func Nbr(a int) int {
	if a == 0 {
		return 0
	}
	i := 0
	for b := 0; b <= a; b++ {
		if b*b == a {
			i = b
			break
		}
	}
	return i
}
func Sqrt(nb int) int {
	i := Nbr(nb)
	return i //this is the base of this func if we want calculat fleat nbr i + (nb-(i^2))/(2*i)
}
func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
