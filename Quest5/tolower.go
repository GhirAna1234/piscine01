package main

import (
	"fmt"
)

func ToLower(s string) string {
	runes := []rune(s)
	for a := 0; a < len(runes); a++ {
		if runes[a] >= 'A' && runes[a] < 'Z' {
			runes[a] = runes[a] + rune(32)
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
