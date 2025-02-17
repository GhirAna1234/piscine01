package main

import (
	"fmt"
)

func BasicAtoi2(s string) int {
	ar := []rune(s)           //covert the string to the rune slice
	n := len(s) - 1           //Stores the length of the string for use in the loop.
	b := 0                    //This variable stores the final converted integer.
	for a := 0; a <= n; a++ { //Iterates through each character of the string.
		if ar[a] < '0' || ar[a] > '9' { //Check if the Character is a Digit.
			return 0
		} else {
			b *= 10
			b += int(ar[a]) - '0' //Convert Character to Integer
		}

	}
	return b //Return the Final Integer
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
