package main

import (
	"fmt"
)

func ExchangeCoin(amount int) []int {
	// Pecahan koin yang tersedia
	pecahan := [...]int{1000, 500, 200, 100, 50, 20, 10, 5, 1}

	// Slice untuk menyimpan hasil kembalian
	result := make([]int, 0)

	// Iterasi melalui setiap pecahan koin
	for _, denom := range pecahan {
		// Jika nilai amount lebih besar dari atau sama dengan denominasi koin saat ini
		for amount >= denom {
			// Tambahkan denominasi koin ke hasil kembalian
			result = append(result, denom)
			// Kurangi nilai amount dengan denominasi koin yang telah ditambahkan
			amount -= denom
		}
	}

	return result
}

func main() {
	fmt.Println(ExchangeCoin(1752)) // [1000, 500, 200, 50, 1, 1]
	fmt.Println(ExchangeCoin(5000)) // [1000, 1000, 1000, 1000, 1000]
	fmt.Println(ExchangeCoin(1234)) // [1000, 200, 20, 10, 1, 1, 1, 1]
	fmt.Println(ExchangeCoin(0))    // []
}
