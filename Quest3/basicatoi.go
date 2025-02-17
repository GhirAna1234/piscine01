package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	v := 0                   //v will store the final integer result.
	for _, char := range s { //Loops through each character (char) in the string.
		//_ ignores the index since we don’t need it.
		v = v*10 + int(char-'0') // Directly convert and accumulate the value
	}
	return v

}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
