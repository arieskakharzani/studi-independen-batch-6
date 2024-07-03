package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {

	numToStr := strconv.Itoa(numbers)

	pair := 0
	maxSum := 0

	for i := 1; i < len(numToStr); i++ {
		number1, _ := strconv.Atoi(string(numToStr[i]))
		number2, _ := strconv.Atoi(string(numToStr[i-1]))

		if number1+number2 > maxSum {
			maxSum = number1 + number2
			pairNumber, _ := strconv.Atoi(string(numToStr[i-1]) + string(numToStr[i]))
			pair = pairNumber
		}
	}
	return pair
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
	fmt.Println(BiggestPairNumber(89083278))
}
