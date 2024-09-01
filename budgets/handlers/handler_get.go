package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Dnreikronos/budgetMannager---Back/budgets/models"
	"github.com/go-chi/chi"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error at trying to parse the id: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	budget, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Error at trying to update the register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(budget)
}
