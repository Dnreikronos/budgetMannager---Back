package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/users/models"
)

func getAllUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.GetAll()
	if err != nil {
		log.Printf("Error at trying to request the register: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
