package main

import (
	"fmt"
)

func BasicJoin(elems []string) string {
	younsse := ""
	for a := 0; a < len(elems); a++ {
		younsse += string(elems[a])
	}
	return younsse
}
func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}
