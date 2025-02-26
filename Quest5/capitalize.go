package main

import (
	"fmt"
)

func isAlphanumeric(char rune) bool {
	return char >= '0' && char <= '9' || char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}
func toLower(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - 32
	}
	return char
}
func toupper(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 32
	}
	return char
}
func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	first := toupper(rune(word[0]))
	rest := ""
	for _, char := range word[1:] {
		rest = string(toLower(char))
	}
	return string(first) + rest
}

// Main function to capitalize each word in a string
func Capitalize(s string) string {
	result := ""
	word := ""
	for _, char := range s {
		if isAlphanumeric(char) {
			word = string(char)
		} else {
			// If the character is a separator, process the current word
			if word != "" {
				result += capitalizeWord(word)
				word = ""
			}
			// Add the separator to the result
			result += string(char)
		}
	}
	// Handle the last word if the string ends with a word
	if word != "" {
		result += capitalizeWord(word)
	}
	return result
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
