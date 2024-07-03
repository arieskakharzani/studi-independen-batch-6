package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	fmt.Print("Jumlah tiket VIP (max 1000) : ")
	fmt.Scan(&VIP)
	fmt.Print("Jumlah tiket REGULER (max 1000) : ")
	fmt.Scan(&regular)
	fmt.Print("Jumlah tiket STUDENT (max 1000) : ")
	fmt.Scan(&student)
	fmt.Print("Masukkan tanggal pemesanan (1-31) : ")
	fmt.Scan(&day)

	jumlahTiket := VIP + regular + student
	totalHarga := (VIP * 30) + (regular * 20) + (student * 10)

	if totalHarga >= 100 {
		if day%2 == 1 {
			if jumlahTiket < 5 {
				return float32(totalHarga) - (float32(totalHarga) * 0.15)
			} else if jumlahTiket >= 5 {
				return float32(totalHarga) - (float32(totalHarga) * 0.25)
			} else {
				return float32(totalHarga)
			}
		} else if day%2 == 0 {
			if jumlahTiket < 5 {
				return float32(totalHarga) - (float32(totalHarga) * 0.1)
			} else if jumlahTiket >= 5 {
				return float32(totalHarga) - (float32(totalHarga) * 0.2)
			} else {
				return float32(totalHarga)
			}
		} else {
			return float32(totalHarga)
		}
	} else {
		return float32(totalHarga)
	}

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
