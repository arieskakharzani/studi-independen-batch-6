package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	operations := strings.Split(calculate, " ")
	if len(operations) == 0 {
		return 0
	}

	number, err := strconv.ParseFloat(operations[0], 32)
	if err != nil {
		return 0
	}

	calculator := internal.NewCalculator(float32(number))
	for i := 1; i < len(operations); i += 2 {
		operation := operations[i]
		num, err := strconv.ParseFloat(operations[i+1], 32)
		if err != nil {
			return 0
		}

		switch operation {
		case "+":
			calculator.Add(float32(num))
		case "-":
			calculator.Subtract(float32(num))
		case "*":
			calculator.Multiply(float32(num))
		case "/":
			calculator.Divide(float32(num))
		default:
			return 0
		}
	}
	return calculator.Result()
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
