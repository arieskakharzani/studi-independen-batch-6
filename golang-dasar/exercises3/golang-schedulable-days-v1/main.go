package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	// Membuat slice untuk menyimpan tanggal-tanggal yang sama-sama kosong
	var availableDates []int
	availableDates = make([]int, 0)

	// Memeriksa setiap tanggal dari slice date1
	for _, d1 := range date1 {
		// Memeriksa setiap tanggal dari slice date2
		for _, d2 := range date2 {
			// Jika tanggal pada date1 sama dengan tanggal pada date2
			if d1 == d2 {
				availableDates = append(availableDates, d1)
				break
			}

		}
	}
	return availableDates
}

func main() {
	fmt.Println(SchedulableDays([]int{11, 12, 13, 14, 15}, []int{5, 10, 12, 13, 20, 21}))
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5}))
	fmt.Println(SchedulableDays([]int{2, 7, 12, 20, 21, 22}, []int{1, 3, 6, 10}))
	fmt.Println(SchedulableDays([]int{}, []int{}))
}
