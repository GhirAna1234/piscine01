package main

import (
	"fmt"
)

func IsUpper(s string) bool {
	for a := 0; a < len(s); a++ {
		if !(s[a] >= 'A' && s[a] <= 'Z') {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}
