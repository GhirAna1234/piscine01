package main

import (
	"fmt"
	"strconv"
)

func atoi(s string) {
	a, e := strconv.Atoi(s)
	fmt.Println(a)
	if e != nil {
		fmt.Println("err")
	}

}
func main() {
	s := "aaa12345"
	atoi(s)
}
