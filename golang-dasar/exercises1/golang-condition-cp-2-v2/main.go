package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	fmt.Print("Input Gender (laki-laki/perempuan) : ")
	fmt.Scan(&gender)
	fmt.Print("Input Tinggi Badan : ")
	fmt.Scan(&height)

	var hasil float64
	if gender == "laki-laki" {
		hasil = float64(height-100) - float64(height-100)*0.1
		return (hasil)
	} else if gender == "perempuan" {
		hasil = float64(height-100) - float64(height-100)*0.15
		return (hasil)
	} else {
		return (hasil)
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
