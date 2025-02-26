package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	nb := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
			nb++
		}
	}
	return nb
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
