package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	}
	availableDates := make([]int, 0)

	// Menyimpan jumlah orang yang akan bersih-bersih pada setiap tanggal
	count := make(map[int]int)

	// Menghitung jumlah orang yang bersedia hadir pada setiap tanggal
	for _, schedule := range villager {
		for _, day := range schedule {
			count[day]++
		}
	}

	// Memeriksa tanggal-tanggal yang memiliki jumlah orang yang sama dengan jumlah total orang
	for day, c := range count {
		if c == len(villager) {
			availableDates = append(availableDates, day)
		}
	}
	return availableDates
}

func main() {
	fmt.Println(SchedulableDays([][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}))
	fmt.Println(SchedulableDays([][]int{{1, 2, 3, 4, 5}, {2, 3, 4, 5}, {2, 3, 4, 10, 11, 12, 15}}))
	fmt.Println(SchedulableDays([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {10, 11, 12}, {21, 22, 23, 24}, {25}}))
	fmt.Println(SchedulableDays([][]int{}))
}
