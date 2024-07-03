package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	fmt.Print("Input Score : ")
	fmt.Scan(&score)
	fmt.Print("Input Absent : ")
	fmt.Scan(&absent)

	if score >= 70 && absent < 5 {
		return ("lulus")
	} else if score < 70 || absent >= 5 {
		return ("tidak lulus")
	} else {
		return ("Tidak Valid")
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(70, 4))
}
