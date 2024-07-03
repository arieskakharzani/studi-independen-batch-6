package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var result []string
	for _, d := range data {
		if strings.Contains(strings.ToLower(d), strings.ToLower(input)) && input != d {
			result = append(result, d)
		}
	}
	return strings.Join(result, ",")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("mobil", "mobil APV", "mobil Avanza", "motor matic", "motor gede"))
	fmt.Println(FindSimilarData("motor", "mobil APV", "mobil Avanza", "motor matic", "motor gede"))
}
