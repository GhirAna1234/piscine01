package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func isvalidbase(base string) bool {
	if len(base) < 2 {
		return false
	}
	charmap := make(map[rune]bool)
	for _, char := range base {
		if charmap[char] || char == '-' || char == '+' {
			return false
		}
		charmap[char] = true
	}
	return true
}
func PrintNbrBase(nbr int, base string) {
	if !isvalidbase(base) {
		fmt.Print("Nv")
		return
	}
	isnegative := false
	if nbr < 0 {
		isnegative = true
		nbr = -nbr
	}
	// convert the number to the given base
	result := ""
	if nbr == 0 {
		result = string(base[0])
	} else {
		for nbr > 0 {
			remainder := nbr % len(base)
			result = string(base[remainder]) + result
			nbr = nbr / len(base)
		}
	}
	if isnegative {
		result = "-" + result
	}
	fmt.Print(result)
}
func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}

// //125
// -1111101
// 7D
// -uoi
// NV
// $
