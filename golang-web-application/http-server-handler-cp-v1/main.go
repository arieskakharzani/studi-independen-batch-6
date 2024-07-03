package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		time := time.Now()

		currentTime := fmt.Sprintf("%s, %d %s %d", time.Weekday(), time.Day(), time.Month(), time.Year())
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(writer, "%s", currentTime)
	}
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
