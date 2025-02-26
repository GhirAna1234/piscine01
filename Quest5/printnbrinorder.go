package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var youn [10]int
	for n > 0 {
		youn[n%10]++
		n /= 10
	}
	for a := 0; a < 10; a++ {
		for youn[a] > 0 {
			z01.PrintRune(rune(a) + '0')
			youn[a]--
		}
	}

}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
