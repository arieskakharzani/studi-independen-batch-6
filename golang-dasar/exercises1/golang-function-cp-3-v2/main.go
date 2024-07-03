package main

import (
	"fmt"
	"strings"
)

// fungsi untuk memisahkan nama berdasarkan simbol
func splitNames(c rune) bool {
	return c == ';' || c == ',' || c == ' '
}

func FindShortestName(names string) string {
	words := strings.FieldsFunc(names, splitNames)
	shortestName := words[0]
	for _, word := range words[1:] {
		if len(word) < len(shortestName) {
			shortestName = word
		} else if len(word) == len(shortestName) && word < shortestName {
			shortestName = word
		}
	}
	return shortestName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
