package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	totalProfits := make(map[int]int)

	for _, branch := range data {
		for month, monthIndex := range branch {
			profit := monthIndex[0] - monthIndex[1]
			totalProfits[month+1] += profit
		}
	}

	numOfMonth := 0
	for key := range totalProfits {
		if key > numOfMonth {
			numOfMonth = key
		}
	}

	finalProfits := make([]int, numOfMonth)
	for key, value := range totalProfits {
		finalProfits[key-1] = value
	}
	return finalProfits
}

func main() {
	fmt.Println(CountProfit([][][2]int{{{1000, 800}, {700, 500}}, {{1000, 800}, {900, 200}}}))

	fmt.Println(CountProfit([][][2]int{{{1000, 500}, {500, 150}, {600, 100}, {800, 750}}}))

	fmt.Println(CountProfit([][][2]int{{{1000, 200}}, {{500, 100}}, {{600, 100}}, {{450, 150}}, {{100, 50}}}))

	fmt.Println(CountProfit([][][2]int{}))
}
