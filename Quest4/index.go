package main

import "fmt"

func Index(s string, toFind string) int {
	count := 0
	for a := 0; a < len(s); a++ {
		if s[a] == toFind {

		}
	}
}
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
