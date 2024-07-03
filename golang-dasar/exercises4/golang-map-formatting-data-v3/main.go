package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	result := make(map[string][]string)

	for _, item := range data {
		parts := strings.Split(item, "-")
		header := parts[0]
		index, _ := strconv.Atoi(parts[1])
		firstOrLast := parts[2]
		value := parts[3]

		if _, ok := result[header]; !ok {
			result[header] = make([]string, 0)
		}

		if firstOrLast == "first" {
			if index >= len(result[header]) {
				result[header] = append(result[header], value)
			} else {
				result[header][index] = value + result[header][index]
			}
		} else {
			if index >= len(result[header]) {
				result[header] = append(result[header], value)
			} else {
				result[header][index] = result[header][index] + " " + value
			}
		}
	}

	return result

}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)
	fmt.Println(res)
	fmt.Println(" ")
	data1 := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar"}
	res1 := ChangeOutput(data1)
	fmt.Println(res1)
}
