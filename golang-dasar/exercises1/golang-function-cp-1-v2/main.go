package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	// Format day to string
	dayStr := fmt.Sprintf("%02d", day)

	// Format month to string
	months := [12]string{
		"January", "February", "March", "April", "May", "June", "July",
		"August", "September", "October", "November", "December",
	}
	monthStr := months[month-1]

	// Format year to string
	yearStr := fmt.Sprintf("%d", year)

	// Gabungkan dalam format yang diinginkan
	return fmt.Sprintf("%s-%s-%s", dayStr, monthStr, yearStr)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
	fmt.Println(DateFormat(31, 12, 2020))
}
