package main

func reverse(s string) string {
	var revString string
	for i := len(s) - 1; i >= 0; i-- {
		revString += string(s[i])
	}
	return revString

}

func isPalindrome(s string) bool {
	revString := reverse(s)
	return revString == s
}

// func main() {
// 	var input string
// 	fmt.Print("Enter a string: ")
// 	fmt.Scanln(&input)

// 	if isPalindrome(input) {
// 		fmt.Println("It is palindrome")
// 	} else {
// 		fmt.Println("It is not palindrome")
// 	}
// }
