package main

import (
	"strings"
	"unicode"
)

func abbreviate(s string) string {
	words := strings.Split(s, " ")
	var abb []string

	for _, word := range words {
		if unicode.IsUpper(rune(word[0])) {
			abb = append(abb, string(word[0]))
		}
	}
	return strings.Join(abb, "")
}

// func main() {
// 	fmt.Println(abbreviate("National Institution of Tecnologies"))
// }
