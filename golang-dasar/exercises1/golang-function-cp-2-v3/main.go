package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowel := "aeiouAEIOU"
	consonant := "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
	vowelCount := 0
	consonantCount := 0
	hasVowel := false
	hasConsonant := false

	for _, char := range str {
		if strings.ContainsRune(vowel, char) {
			vowelCount++
			hasVowel = true
		}
		if strings.ContainsRune(consonant, char) {
			consonantCount++
			hasConsonant = true
		}
		if unicode.IsSpace(char) {
			continue
		}
	}

	return vowelCount, consonantCount, !(hasVowel && hasConsonant)
}

// gunakan untuk lakukan debug
func main() {
	fmt.Println(CountVowelConsonant("kopi"))
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
