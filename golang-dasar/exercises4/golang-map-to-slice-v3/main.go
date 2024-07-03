package main

import (
	"fmt"
)

func MapToSlice(mapData map[string]string) [][]string {
	result := make([][]string, 0)

	for key, value := range mapData {
		result = append(result, []string{key, value})
	}

	return result
}

func main() {
	mapData := map[string]string{"hello": "world", "John": "Doe", "age": "14"}
	output := MapToSlice(mapData)
	fmt.Println(output)

	mapData = map[string]string{"foo": "33", "bar": "44", "baz": "55"}
	output = MapToSlice(mapData)
	fmt.Println(output)

	mapData = map[string]string{}
	output = MapToSlice(mapData)
	fmt.Println(output)
}
