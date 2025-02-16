package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	arr := []rune{}
	n := len(s) - 1
	v := 0
	for a := 0; a <= n; a++ {
		arr = append(arr, rune(s[a]-48))
	}

	for b := 0; b <= len(arr)-1; b++ {
		v = v*10 + int(arr[b])

	}
	return v
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
