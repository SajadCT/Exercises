package main

import (
	"strings"
)

func panagram(s string) bool {
	s = strings.ToLower(s)
	letters := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			letters[string(ch)] = true
		}
	}
	count := len(letters)
	return count == 26
}

// func main() {
// 	fmt.Println(panagram("hi friends how are you"))
// }
