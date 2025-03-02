package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	for _, y := range args {
		if y == "--upper" {
			upper = true
			args = args[1:]
		}
	}
	for _, arg := range args {
		count := 0
		for _, i := range arg {
			count = count*10 + int(i-'0')
		}
		if count >= 1 && count <= 26 {
			if upper == false {
				z01.PrintRune(rune(count + 96))
			} else {
				z01.PrintRune(rune(count + 64))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
