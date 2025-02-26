package main

import (
	"fmt"
)

func ToUpper(s string) string {
	runes := []rune(s)
	for a := 0; a < len(runes); a++ {
		if runes[a] >= 'a' && runes[a] <= 'z' {
			runes[a] = runes[a] - rune(32)
		}
	}
	return string(runes)
}
func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
