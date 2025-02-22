package main

import (
	"fmt"
)

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
func FindNextPrime(nb int) int {
	for !IsPrime(nb) {
		nb++
	}
	return nb
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}
