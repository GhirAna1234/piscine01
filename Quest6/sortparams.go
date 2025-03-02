package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for char := 0; char < len(args); char++ {
		for char2 := 0; char2 < len(args)-char-1; char2++ {
			if args[char2] > args[char+1] {
				args[char2], args[char2+1] = args[char2+1], args[char2]
			}
		}
	}
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
