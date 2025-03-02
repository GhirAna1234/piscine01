package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	isupper := false
	for _, r := range args {
		if r == "--upper" {
			isupper = true
			args = args[1:]
		}
	}
	for _, char := range args {
		count := 0
		for _, a := range char {
			count = count*10 + int(a-'0')
		}
		if !(count >= 1 && count <= 26) {
			z01.PrintRune(' ')
		}
		if !isupper && (count >= 1 && count <= 26) {
			z01.PrintRune(rune(count + 64))
		} else {
			z01.PrintRune(rune(count + 96))
		}

	}
	z01.PrintRune('\n')
}
