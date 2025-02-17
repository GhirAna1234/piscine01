package main

import (
	"github.com/01-edu/z01"
)

func printalphabet() {
	for a := 'a'; a <= 'z'; a++ {
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}
func main() {
	printalphabet()
}
