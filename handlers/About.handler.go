package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type aboutResponse struct {
	Title string `json:"title"`
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("bearer")

	fmt.Println("bearer Token: ", token)
	about := aboutResponse{
		Title: "About",
	}
	err := json.NewEncoder(w).Encode(about)

	if err != nil {
		println(err)
	}
}
