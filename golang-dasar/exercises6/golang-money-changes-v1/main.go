package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	totalPrice := 0

	for _, product := range products {
		totalPrice += product.Price + product.Tax
	}
	change := amount - totalPrice

	if amount == totalPrice {
		return []int{}
	}

	if change < 0 {
		return nil
	}

	var result []int
	for _, v := range []int{1000, 500, 200, 100, 50, 20, 10, 5, 1} {
		for change >= v {
			change -= v
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println(MoneyChanges(10000, []Product{{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300}}))
	fmt.Println(MoneyChanges(30000, []Product{{Name: "baju 1", Price: 10000, Tax: 1000}, {Name: "Sepatu", Price: 15550, Tax: 1555}}))
	fmt.Println(MoneyChanges(5500, []Product{{Name: "Baju", Price: 5000, Tax: 500}}))
	fmt.Println(MoneyChanges(0, []Product{}))
}
