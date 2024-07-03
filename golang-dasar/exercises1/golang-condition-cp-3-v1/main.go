package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	fmt.Print("Input Nilai Matematika (1-100) : ")
	fmt.Scan(&math)
	fmt.Print("Input Nilai Sains (1-100) : ")
	fmt.Scan(&science)
	fmt.Print("Input Nilai Bahasa Inggris (1-100) : ")
	fmt.Scan(&english)
	fmt.Print("Input Nilai Bahasa Indonesia (1-100) : ")
	fmt.Scan(&indonesia)

	avg := (math + science + english + indonesia) / 4
	if avg == 100 {
		return ("Sempurna")
	} else if avg >= 90 {
		return ("Sangat Baik")
	} else if avg >= 80 {
		return ("Baik")
	} else if avg >= 70 {
		return ("Cukup")
	} else if avg >= 60 {
		return ("Kurang")
	} else {
		return ("Sangat Kurang")
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
