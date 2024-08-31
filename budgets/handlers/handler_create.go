package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/budgets/models"
)

func createBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget

	err := json.NewDecoder(r.Body).Decode(&budget)
	if err != nil {
		log.Printf("Error in trying do to decode of Json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.InsertBudget(budget)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("An error occored when trying to insert: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("User inserted with sucess! ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
