package main

import "fmt"

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}

	average := 0.0
	min := s.Grades[0]
	max := s.Grades[0]

	for _, grade := range s.Grades {
		average += float64(grade)
		if grade < min {
			min = grade
		}
		if grade > max {
			max = grade
		}
	}

	average /= float64(len(s.Grades))

	return average, min, max
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 100, 90, 100, 90},
	})

	fmt.Println(avg, min, max)
}
