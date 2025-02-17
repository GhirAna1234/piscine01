package main

import "github.com/01-edu/z01"

func printdigits() {
	for a := '0'; a <= '9'; a++ {
		z01.PrintRune((rune(a)))
	}
	z01.PrintRune('\n')
}
func main() {
	printdigits()
}
