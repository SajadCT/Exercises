package main
import (
    "fmt"
    "strings"
)
func frequency(s string) map[string]int { 
	   letterFreq := make(map[string]int)    s = strings.ToLower(s)    
	   for i := 0; i < len(s); i++ {
        ch := s[i]
        if ch >= 'a' && ch <= 'z' {
            letterFreq string(ch)] = letterFreq string(ch)] + 1
        }
    }    return letterFreq
}


func main() {
    fmt.Println(frequency("tomorrow"))
}