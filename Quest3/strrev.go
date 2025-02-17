package main

import "fmt"

func StrRev(s string) string {
	runes := []rune(s)                           // Convert String to Rune Slice
	r := len(runes)                              //Stores the length of the slice in r.
	for a, b := 0, r-1; a < b; a, b = a+1, b-1 { //Swap Characters to Reverse the String
		runes[a], runes[b] = runes[b], runes[a]
	}
	return string(runes) //Converts the modified rune slice back into a string.
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
