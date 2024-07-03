package main

import (
	"fmt"
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	studentHandler := http.HandlerFunc(StudentHandler())
	http.Handle("/student", RequestMethodGet(studentHandler))
	fmt.Println("Server is running on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
