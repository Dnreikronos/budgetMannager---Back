package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/bills/models"
)

func CreateBills(w http.ResponseWriter, r *http.Request) {
	var bills models.Bills

	err := json.NewDecoder(r.Body).Decode(&bills)
	if err != nil {
		log.Printf("Error in tying to decode of json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.CreateBills(bills)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("An error occorred when tying to insert: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Bill inserted with sucess! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
