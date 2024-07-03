package main

import "fmt"

func ReverseData(arr [5]int) [5]int {
	// Inisialisasi array baru untuk menyimpan hasil yang dibalik
	var reversed [5]int

	// Perulangan melalui array input
	for i, num := range arr {
		// Inisialisasi variabel untuk menyimpan hasil pembalikan setiap angka
		var temp int = 0

		// Proses pembalikan setiap angka
		for num > 0 {
			digit := num % 10
			temp = temp*10 + digit
			num /= 10
		}

		// Memasukkan angka yang telah dibalik ke dalam array hasil
		reversed[4-i] = temp
	}

	return reversed
}

func main() {
	fmt.Println(ReverseData([5]int{123, 456, 11, 1, 2}))
	fmt.Println(ReverseData([5]int{456789, 44332, 2221, 12, 10}))
	fmt.Println(ReverseData([5]int{10, 10, 10, 10, 10}))
	fmt.Println(ReverseData([5]int{23456, 789, 123, 456, 500}))
	fmt.Println(ReverseData([5]int{0, 0, 0, 0, 0}))

}
