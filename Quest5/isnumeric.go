package main

import (
	"fmt"
)

func IsNumeric(s string) bool {
	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
