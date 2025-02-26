package main

import (
	"fmt"
)

func IsLower(s string) bool {
	for _, a := range s {
		if !(a >= 'a' && a <= 'z') {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}
