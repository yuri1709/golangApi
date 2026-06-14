package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// helper para responder JSON
func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			ID:   "1",
			Name: "Yuri Roliz",
		},
	}
	respondJSON(w, http.StatusOK, users)
}
