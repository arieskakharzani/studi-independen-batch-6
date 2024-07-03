package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type AIModelConnector struct {
	Client *http.Client
}

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type Response struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

func CsvToSlice(data string) (map[string][]string, error) {
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, errors.New("CSV file must contain at least one row of data")
	}

	headers := records[0]
	result := make(map[string][]string)
	for _, header := range headers {
		result[header] = []string{}
	}

	for _, row := range records[1:] {
		for i, value := range row {
			result[headers[i]] = append(result[headers[i]], value)
		}
	}

	return result, nil
}

func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	url := "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq"
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return Response{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	fmt.Printf("Response body: %s\n", string(respBody))

	var response Response
	err = json.NewDecoder(bytes.NewReader(respBody)).Decode(&response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func recommendEnergySaving(response Response, dataMap map[string][]string) {
	if response.Answer != "" {
		fmt.Println("Rekomendasi Penghematan Energi:")
		fmt.Println("1. Matikan alat yang tidak digunakan.")
		fmt.Println("2. Gunakan peralatan dengan konsumsi energi lebih rendah.")

		// Tambahkan rekomendasi berdasarkan data alat dan konsumsi energi
		consumptionMap := make(map[string]float64)
		for i, appliance := range dataMap["Appliance"] {
			if status := dataMap["Status"][i]; status == "On" {
				consumption, err := strconv.ParseFloat(dataMap["Energy_Consumption"][i], 64)
				if err == nil {
					consumptionMap[appliance] += consumption
				}
			}
		}

		fmt.Println("3. Rekomendasi berdasarkan konsumsi energi alat:")
		for appliance, consumption := range consumptionMap {
			fmt.Printf("   - %s: Total konsumsi %.2f kWh. Pertimbangkan untuk mengurangi penggunaan.\n", appliance, consumption)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	fileData, err := ioutil.ReadFile("data-series.csv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	dataMap, err := CsvToSlice(string(fileData))
	if err != nil {
		fmt.Println("Error converting CSV to map:", err)
		return
	}

	query := "What is the total energy consumption for the next day?"
	payload := Inputs{
		Table: dataMap,
		Query: query,
	}

	fmt.Printf("Payload: %+v\n", payload)

	client := &http.Client{}
	connector := AIModelConnector{Client: client}
	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		fmt.Println("Error: Hugging Face token is not set")
		return
	}

	response, err := connector.ConnectAIModel(payload, token)
	if err != nil {
		fmt.Println("Error connecting to AI model:", err)
		return
	}

	fmt.Println("Predicted Answer:", response.Answer)
	fmt.Println("Coordinates:", response.Coordinates)
	fmt.Println("Cells:", response.Cells)
	fmt.Println("Aggregator:", response.Aggregator)

	recommendEnergySaving(response, dataMap)
}
