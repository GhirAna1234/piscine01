package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {

	name := os.Args[1:]
	for _, i := range name {
		for _, a := range i {
			z01.PrintRune(rune(a))
		}
		z01.PrintRune('\n')
	}

}
