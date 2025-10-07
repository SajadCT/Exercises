package main

import (
	"fmt"
	"strings"
)

func panagram(s string) bool {
	s = strings.ToLower(s)
	checkLetters := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			checkLetters[string(ch)] = true
		}
	}
	count := len(checkLetters)
	return count == 26
}

func main() {
	fmt.Println(panagram("hi iam sajad from"))
}
