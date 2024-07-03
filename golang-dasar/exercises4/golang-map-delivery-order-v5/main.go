package main

import (
	"fmt"
	"math"
	"strings"
)

func DeliveryOrder(data []string, day string) map[string]float32 {
	result := make(map[string]float32)

	// Tentukan biaya admin berdasarkan hari
	var adminFee float32
	if day == "senin" || day == "rabu" || day == "jumat" {
		adminFee = 0.1
	} else {
		adminFee = 0.05
	}

	// Tentukan kota yang dapat dikirim pada hari tertentu
	var validLocations map[string]bool
	switch day {
	case "senin":
		validLocations = map[string]bool{"JKT": true, "DPK": true}
	case "selasa":
		validLocations = map[string]bool{"JKT": true, "BKS": true, "DPK": true}
	case "rabu":
		validLocations = map[string]bool{"BDG": true}
	case "kamis":
		validLocations = map[string]bool{"BDG": true, "BKS": true}
	case "jumat":
		validLocations = map[string]bool{"BKS": true}
	case "sabtu":
		validLocations = map[string]bool{"JKT": true, "BDG": true}
	}

	// Perulangan untuk data pesanan
	for _, order := range data {
		orderDetails := strings.Split(order, ":")
		firstName := orderDetails[0]
		lastName := orderDetails[1]
		price := orderDetails[2]
		location := orderDetails[3]

		// Cek apakah lokasi dapat dikirim pada hari tertentu
		if _, ok := validLocations[location]; ok {
			totalPrice := parseFloat(price)
			totalWithFee := totalPrice * (1 + adminFee)
			// Bulatkan total biaya pengiriman ke 2 desimal
			totalWithFee = float32(math.Round(float64(totalWithFee)*100) / 100)
			result[firstName+"-"+lastName] = totalWithFee
		}
	}
	return result
}

func parseFloat(s string) float32 {
	var result float32
	_, err := fmt.Sscanf(s, "%f", &result)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}
	day := "sabtu"
	deliveryData := DeliveryOrder(data, day)
	fmt.Println("Test Case 1:")
	printMap(deliveryData)

	fmt.Println("\nTest Case 2:")
	data1 := []string{
		"Anggi:Anggraini:20000:DPK",
		"Andi:Sukirman:15000:JKT",
	}
	day1 := "selasa"
	deliveryData1 := DeliveryOrder(data1, day1)
	printMap(deliveryData1)
}

func printMap(data map[string]float32) {
	for key, value := range data {
		fmt.Printf("%s: %.2f\n", key, value)
	}
}
