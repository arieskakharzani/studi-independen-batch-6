package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var report Report
	err = json.Unmarshal(file, &report)
	if err != nil {
		panic(err)
	}

	return report, nil
}

func GradePoint(report Report) float64 {
	TotalScore := 0.0
	TotalCredit := 0.0

	for _, study := range report.Studies {
		switch study.Grade {
		case "A":
			TotalScore += 4.0 * float64(study.StudyCredit)
		case "AB":
			TotalScore += 3.5 * float64(study.StudyCredit)
		case "B":
			TotalScore += 3.0 * float64(study.StudyCredit)
		case "BC":
			TotalScore += 2.5 * float64(study.StudyCredit)
		case "C":
			TotalScore += 2.0 * float64(study.StudyCredit)
		case "CD":
			TotalScore += 1.5 * float64(study.StudyCredit)
		case "D":
			TotalScore += 1.0 * float64(study.StudyCredit)
		case "DE":
			TotalScore += 0.5 * float64(study.StudyCredit)
		case "E":
			TotalScore += 0.0 * float64(study.StudyCredit)
		default:
			return 0.0
		}
		TotalCredit += float64(study.StudyCredit)
	}

	if TotalCredit == 0 {
		return 0.0
	}

	return TotalScore / TotalCredit
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
