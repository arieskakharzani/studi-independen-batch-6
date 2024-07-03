package main

import "fmt"

func Sortheight(height []int) []int {
	// Jika slice height kosong, kembalikan slice kosong
	if len(height) == 0 {
		return height
	}

	// Menggunakan algoritma pengurutan bubble sort untuk mengurutkan tinggi badan
	for i := 0; i < len(height)-1; i++ {
		for j := 0; j < len(height)-i-1; j++ {
			if height[j] > height[j+1] {
				// Tukar posisi elemen jika elemen sekarang lebih besar dari elemen berikutnya
				height[j], height[j+1] = height[j+1], height[j]
			}
		}
	}
	return height
}

func main() {
	fmt.Println(Sortheight([]int{172, 170, 150, 165, 144, 155, 159}))
	fmt.Println(Sortheight([]int{155, 156, 160, 161, 162, 165, 170, 172}))
	fmt.Println(Sortheight([]int{180, 177, 175, 173, 170, 166, 165, 160}))
	fmt.Println(Sortheight([]int{}))
}
