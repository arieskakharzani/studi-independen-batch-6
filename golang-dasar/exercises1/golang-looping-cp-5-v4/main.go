package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	// Memisahkan string menjadi kata-kata
	words := strings.Fields(str) 
	reversed := make([]string, len(words))

	// Mengecek apakah semua kata dalam string adalah huruf kapital
	allCapital := true
	for _, word := range words {
		for _, char := range word {
			if !unicode.IsUpper(char) {
				allCapital = false
				break
			}
		}
		if !allCapital {
			break
		}
	}

	for i, word := range words {
		// Mengecek apakah huruf pertama dalam kata adalah huruf kapital
		firstLetterIsCapital := unicode.IsUpper(rune(word[0]))

		// Mengubah kata menjadi huruf kecil agar dapat dibalik
		word = strings.ToLower(word)

		// Membalik kata
		reversedWord := ""
		for _, char := range word {
			reversedWord = string(char) + reversedWord
		}

		// Jika semua kata dalam string adalah huruf kapital, maka setelah dibalik, tetapkan huruf kapital pada setiap kata
		if allCapital {
			reversedWord = strings.ToUpper(reversedWord)
		} else if firstLetterIsCapital { // Jika huruf pertama dalam kata adalah huruf kapital dan tidak semua huruf dalam kata adalah kapital
			reversedWord = strings.ToUpper(string(reversedWord[0])) + reversedWord[1:]
		}

		reversed[i] = reversedWord
	}

	// Menggabungkan kata-kata yang telah dibalik
	return strings.Join(reversed, " ")
}

func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
	fmt.Println(ReverseWord("ini terlalu mudah"))
	fmt.Println(ReverseWord("KITA SELALU BERSAMA"))
}
