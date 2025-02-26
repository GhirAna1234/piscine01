package main

import (
	"fmt"
)

// Helper function to check if a character is a digit
func isDigit(char rune) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

// Function to extract digits and optional signs from the string
func extractDigits(s string) (string, int) {
	digits := ""
	foundDigits := false
	sign := 1
	for _, char := range s {
		// Handle optional sign
		if (char == '-' || char == '+') && !foundDigits {
			if char == '-' {
				sign = -1
			}
		}
		// Extract digits
		if isDigit(char) {
			digits += string(char)
			foundDigits = true
		}
	}
	return digits, sign
}

// Main function to convert the extracted digits into an integer
func TrimAtoi(s string) int {
	count := 0
	digits, sign := extractDigits(s)
	for _, char := range digits {
		count = count*10 + int(char-'0')
	}
	return count * sign
}
func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
