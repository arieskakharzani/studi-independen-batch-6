package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	if len(data) == 0 {
		return []string{}, nil
	}

	return data, nil
}

func CalculateProfitLoss(data []string) string {
	totalIncome := 0
	totalExpense := 0
	lastDate := ""

	for _, transaction := range data {
		if transaction == "" {
			continue
		}
		transactionData := strings.Split(transaction, ";")
		// Mendapatkan tanggal, tipe transaksi, dan jumlah
		date := transactionData[0]
		transactionType := transactionData[1]
		amount, _ := strconv.Atoi(transactionData[2])

		// Memperbarui tanggal terakhir
		if date > lastDate {
			lastDate = date
		}

		// Menghitung total income dan expense
		if transactionType == "income" {
			totalIncome += amount
		} else if transactionType == "expense" {
			totalExpense += amount
		}
	}

	// Menghitung profit atau loss
	profitLoss := totalIncome - totalExpense

	// Menentukan tipe profit atau loss
	profitLossType := "profit"
	if profitLoss < 0 {
		profitLossType = "loss"
		profitLoss *= -1
	}

	// Memformat hasil dan mengembalikan nilai
	return fmt.Sprintf("%s;%s;%d", lastDate, profitLossType, profitLoss)
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
