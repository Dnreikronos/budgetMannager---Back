package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	models "github.com/Dnreikronos/budgetMannager---Back/models/users"
)

func getAllUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.GetAll()
	if err != nil {
		log.Println("Error at trying to request the register: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
