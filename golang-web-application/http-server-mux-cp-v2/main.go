package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time := time.Now()
		currentTime := fmt.Sprintf("%s, %d %s %d", time.Weekday(), time.Day(), time.Month(), time.Year())
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "%s", currentTime)
	}
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintf(w, "Hello there")
		} else {
			fmt.Fprintf(w, "Hello, %s!", name)
		}
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
	fmt.Println("Server is running on localhost:8080")
}
