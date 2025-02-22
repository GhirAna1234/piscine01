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
func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}
