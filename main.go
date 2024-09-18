package main

import (
	"fmt"
	"golang/auth"
	"golang/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.Handle("/about", auth.ProtectedRoute(http.HandlerFunc(handlers.AboutHandler)))
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Print("Error Starting Server:", err)
	}
}
