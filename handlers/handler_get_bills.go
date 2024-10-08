package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	querys "github.com/Dnreikronos/budgetMannager---Back/querys/bills"
	"github.com/go-chi/chi"
)

func GetBills(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error at trying to parse ID: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	bills, err := querys.GetBills(int64(id))
	if err != nil {
		log.Printf("Error at trying to update register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bills)
}
