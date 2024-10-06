package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	querys "github.com/Dnreikronos/budgetMannager---Back/querys/bills"
)

func GetAllBills(w http.ResponseWriter, r *http.Request) {
	bills, err := querys.GetAll()
	if err != nil {
		log.Printf("Error at trying to parse the register: %v", err)
	}

	w.Header().Add("Content-Type", "apllication/json")
	json.NewEncoder(w).Encode(bills)
}
