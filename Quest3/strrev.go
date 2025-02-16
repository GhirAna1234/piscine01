package main

import "fmt"

func StrRev(s string) string {
	runes := []rune(s)
	r := len(runes)
	for a, b := 0, r-1; a < b; a, b = a+1, b-1 {
		runes[a], runes[b] = runes[b], runes[a]
	}
	return string(runes)
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
