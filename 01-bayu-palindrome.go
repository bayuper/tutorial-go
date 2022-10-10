package main

import "fmt"


func isPalindrome(str string) string {
	reverseStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}

	for i := range(str) {
		if str[i] != reverseStr[i] {
			return "Not Palindrome"
		}
		
	}
	return "Palindrome"
}

func main() {
	var str string
	
	fmt.Println("Input word: ")
	fmt.Scanf("%s", &str)

	fmt.Printf("%v", isPalindrome(str))
}