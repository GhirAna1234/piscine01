package main

import (
	"fmt"
)

func UltimateDivMod(a *int, b *int) {

	c := *a
	*a = c / *b
	*b = c % *b

}
func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println("this is the address of a:", &a)
	fmt.Println("this is the address of b:", &b)
	fmt.Println(a)
	fmt.Println(b)
}
