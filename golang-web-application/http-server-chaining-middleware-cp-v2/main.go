package main

import (
	"fmt"
	"net/http"
)

func AdminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Admin page"))
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("role")
		if role != "ADMIN" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Role not authorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	adminHandler := http.HandlerFunc(AdminHandler())
	http.Handle("/admin", RequestMethodGetMiddleware(AdminMiddleware(adminHandler)))
	fmt.Println("Server is running on localhost:8080")

	http.ListenAndServe("localhost:8080", nil)
}
