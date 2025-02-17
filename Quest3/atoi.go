package main

import (
	"fmt"
)

func zb(s string) int {
	v := 0
	for _, char := range s {
		v = v*10 + int(char-'0')
	}
	return v
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' && s[1] == '-' || s[0] == '+' && s[1] == '+' {
		return 0
	}
	for i := 0; i < len(s)-1; i++ {
		if !(s[i] >= '0' && s[i] <= '9' || s[i] == '+' || s[i] == '-') {
			return 0
		}
	}
	c := 1
	v := 0
	if s[0] == '-' {
		c = -1
		s = s[1:]
	}
	if s[0] == '+' {
		s = s[1:]
	}
	v = zb(s)
	return v * c

}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
