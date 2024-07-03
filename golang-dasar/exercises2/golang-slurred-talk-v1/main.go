package main

import "fmt"

func SlurredTalk(words *string) {
	changeLetter := ""

	for _, letter := range *words {
		if letter == 'S' || letter == 'R' || letter == 'Z' {
			changeLetter += "L"
		} else if letter == 's' || letter == 'r' || letter == 'z' {
			changeLetter += "l"
		} else {
			changeLetter += string(letter)
		}
	}

	*words = changeLetter
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = ""
	SlurredTalk(&words)
	fmt.Println(words)

	words = "Steven"
	SlurredTalk(&words)
	fmt.Println(words) // output: Lteven

	words = "Saya Steven"
	SlurredTalk(&words)
	fmt.Println(words) // output: Laya Lteven

	words = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
