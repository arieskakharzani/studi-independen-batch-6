package main

import (
	"fmt"
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	// TODO: replace this
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}

		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
			return
		}

		if IsNameExists(name) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Name is exists"))
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
		}
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe("localhost:8080", GetMux())
}
