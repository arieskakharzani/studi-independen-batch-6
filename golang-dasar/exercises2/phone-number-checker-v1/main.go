package main

import (
	"fmt"
	"strings"
)

func PhoneNumberChecker(number string, result *string) {
	if number == "" {
		*result = "invalid"
		return
	}

	if number[0:2] == "08" {
		number = "62" + number[1:]
	}

	prefix := number[0:3]
	if prefix != "628" {
		*result = "invalid"
		return
	}

	if len(number) < 11 {
		*result = "invalid"
		return
	}

	if strings.HasPrefix(number, "62811") || strings.HasPrefix(number, "62812") || strings.HasPrefix(number, "62813") || strings.HasPrefix(number, "62814") || strings.HasPrefix(number, "62815") {
		*result = "Telkomsel"
	} else if strings.HasPrefix(number, "62816") || strings.HasPrefix(number, "62817") || strings.HasPrefix(number, "62818") || strings.HasPrefix(number, "62819") {
		*result = "Indosat"
	} else if strings.HasPrefix(number, "62821") || strings.HasPrefix(number, "62822") || strings.HasPrefix(number, "62823") {
		*result = "XL"
	} else if strings.HasPrefix(number, "62827") || strings.HasPrefix(number, "62828") || strings.HasPrefix(number, "62829") {
		*result = "Tri"
	} else if strings.HasPrefix(number, "62852") || strings.HasPrefix(number, "62853") {
		*result = "AS"
	} else if strings.HasPrefix(number, "62881") || strings.HasPrefix(number, "62882") || strings.HasPrefix(number, "62883") || strings.HasPrefix(number, "62884") || strings.HasPrefix(number, "62885") || strings.HasPrefix(number, "62886") || strings.HasPrefix(number, "62887") || strings.HasPrefix(number, "62888") {
		*result = "Smartfren"
	} else {
		*result = "invalid"
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "6281111111111"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)

	number = "08193456123"
	PhoneNumberChecker(number, &result)
	fmt.Println(result)

	number = "628523456789"
	PhoneNumberChecker(number, &result)
	fmt.Println(result)

	number = "081234567"
	PhoneNumberChecker(number, &result)
	fmt.Println(result)

	number = "08222"
	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
