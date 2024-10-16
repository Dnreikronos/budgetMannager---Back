package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dnreikronos/budgetMannager---Back/models"
	querys "github.com/Dnreikronos/budgetMannager---Back/querys/bills"
)

func DeleteBills(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
	}

	bills := models.Bills{ID: id}

	deletedID, err := querys.DeleteBills(bills)
	if err != nil {
		http.Error(w, "Error deleting bills", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Bills with ID %d has been deleted", deletedID),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error on deleting response", http.StatusInternalServerError)
	}
}
