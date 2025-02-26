package main

import (
	"fmt"
)

func IsPrintable(s string) bool {
	for _, v := range s {
		if !(v >= ' ' && v <= '~') {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}
