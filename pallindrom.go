package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// Convert to lowercase to make it case-insensitive
	s = strings.ToLower(s)

	// Compare characters from both ends
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
