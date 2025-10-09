package main

func frequency(s string) map[string]int {
	letterFreq := make(map[string]int)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		letterFreq[string(ch)] = letterFreq[string(ch)] + 1
	}
	return letterFreq
}

// func main() {
// 	fmt.Println(frequency("Are u happy ?"))
// }
