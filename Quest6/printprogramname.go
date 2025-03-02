package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	word := ""
	sr := string(os.Args[0])
	for a := 0; a < len(sr); a++ {
		if sr[a] == '/' {
			word = ""
		} else {
			word += string(sr[a])
		}

	}
	for a := range word {
		z01.PrintRune(rune(word[a]))
	}
	z01.PrintRune('\n')
}
