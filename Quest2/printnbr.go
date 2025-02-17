package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {

	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	l := 0
	arr := []int{}
	for n > 0 {
		l = n % 10
		arr = append(arr, l)
		n = n / 10

	}
	for i := len(arr) - 1; i >= 0; i-- {
		z01.PrintRune(rune(arr[i]) + '0')

	}

}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
