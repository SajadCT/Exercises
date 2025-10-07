package exercises
package main
import (
    "fmt"
    "strings"
    "unicode"
)
func Abbri(s string) string {  
    words := strings.Split(s, " ")
    var abbriviate []string    
	
	for _,word := range words 
	{
        if unicode.IsUpper(rune(word[0])) {
            abbriviate = append(abbriviate, string(word[0]))
        }
    }   
	 return strings.Join(abbriviate, "")}
	func main() 
	{   
		 fmt.Println(Abbri("hai i am sajad from kozhikode"))
}