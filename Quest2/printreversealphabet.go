package main

import "github.com/01-edu/z01"

func printreversealphabet() {
	for a := 'z'; a >= 'a'; a-- {
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}
func main() {
	printreversealphabet()
}
