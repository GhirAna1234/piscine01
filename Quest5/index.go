package main

import (
	"fmt"
)

func Index(s string, toFind string) int {
	for a := 0; a < len(s)-len(toFind); a++ {
		if s[a:a+len(toFind)] == toFind {
			return a
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
