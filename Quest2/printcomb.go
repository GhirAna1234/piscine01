package main

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		for b := '1'; b <= '8'; b++ {
			for c := '2'; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a != '7' || b != '8' || c != '9' {
					z01.PrintRune(',')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
func main() {
	PrintComb()
}
