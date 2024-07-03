package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	count := 0

	for i := 0; i < len(text); i++ {
		if text[i] == 'R' || text[i] == 'r' || text[i] == 'S' || text[i] == 's' || text[i] == 'T' || text[i] == 't' || text[i] == 'Z' || text[i] == 'z' {
			count++
		}
	}
	return count

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
	fmt.Println(CountingLetter("Remaja muda yang berbakat"))
	fmt.Println(CountingLetter("Zebra Zig Zag"))
}
