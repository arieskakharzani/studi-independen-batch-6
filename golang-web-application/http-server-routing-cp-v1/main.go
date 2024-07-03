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
		if name != "" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, " + name + "!"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello there"))
		}
	}
}

func main() {
	// TODO: answer here
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())

	http.ListenAndServe("localhost:8080", nil)
}
