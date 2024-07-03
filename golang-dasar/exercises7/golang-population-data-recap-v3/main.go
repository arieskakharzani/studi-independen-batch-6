package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	result := make([]map[string]interface{}, len(data))

	for i, d := range data {
		personData := strings.Split(d, ";")

		person := make(map[string]interface{})
		person["name"] = personData[0]
		age, _ := strconv.Atoi(personData[1])
		person["age"] = age
		person["address"] = personData[2]

		if len(personData[3]) > 0 {
			height, _ := strconv.ParseFloat(personData[3], 64)
			person["height"] = height
		}

		if len(personData[4]) > 0 {
			isMarried, _ := strconv.ParseBool(personData[4])
			person["isMarried"] = isMarried
		}

		result[i] = person
	}
	return result
}

func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	fmt.Println(PopulationData(data))

	data2 := []string{"Jaka;25;Jakarta;false;170.1", "Anggi;24;Bandung;;"}
	fmt.Println(PopulationData(data2))

	data3 := []string{}
	fmt.Println(PopulationData(data3))

	data4 := []string{"Budi;23;Jakarta;170.1;true"}
	fmt.Println(PopulationData(data4))
}
