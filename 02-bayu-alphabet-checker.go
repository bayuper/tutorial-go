package main

import (
	"fmt"
	"unicode"
)

func isVowel(input rune) string {
	var char rune = unicode.ToLower(input)
	if (char == 'a') || (char == 'i') || (char == 'u') || (char == 'e') || (char == 'o') {
		return "Vowel"
	} else{
		return "Consonant"
	}
}

func main() {
	var char rune

	fmt.Printf("Input char: ")
	fmt.Scanf("%c\n", &char)

	fmt.Printf("%s", isVowel(char))
}