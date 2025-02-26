package main

import (
	"fmt"
)

func IsAlpha(s string) bool {
	for a := 0; a < len(s); a++ {
		if !((s[a] >= 'a' && s[a] <= 'z') || (s[a] >= 'A' && s[a] <= 'Z')) || (s[a] >= '0' && s[a] <= '9') {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}
