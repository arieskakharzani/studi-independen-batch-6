package main

import "fmt"

func TicketPlayground(height, age int) int {

	harga := 15000

	if age < 5 {
		return -1
	}
	if age >= 5 || height > 120 {
		harga += 0
	}
	if age >= 8 || height > 135 {
		harga += 10000
	}
	if age >= 10 || height > 150 {
		harga += 15000
	}
	if age == 12 || height > 160 {
		harga += 20000
	}
	if age > 12 {
		harga = 100000
	}
	return harga

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
