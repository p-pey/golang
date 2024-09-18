package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
	w.Header().Set("Content-Type", "application/json")
	response := HomeResponse{
		Message: "Hello from home",
		Status:  200,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("Home Page: ", err)
	}
}
