package main

import (
	"fmt"
)

func Join(strs []string, sep string) string {
	younss := ""
	sep = ":"
	for a := 0; a < len(strs); a++ {
		younss += string(strs[a])
		if a != len(strs)-1 {
			younss += sep
		}
	}
	return younss
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
