package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func getAllUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.getAll()
	if err != nil {
		log.Println("Error at trying to request the register: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
