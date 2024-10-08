package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	querys "github.com/Dnreikronos/budgetMannager---Back/querys/budgets"
)

func GetAllBudget(w http.ResponseWriter, r *http.Request) {
	budget, err := querys.GetAll()
	if err != nil {
		log.Printf("Error at trying to request the register: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(budget)
}
