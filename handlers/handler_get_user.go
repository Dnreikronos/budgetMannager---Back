package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	querys "github.com/Dnreikronos/budgetMannager---Back/querys/user"
	"github.com/go-chi/chi"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error at trying to do parse of id: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := querys.Get(int64(id))
	if err != nil {
		log.Printf("Error at trying to update register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
