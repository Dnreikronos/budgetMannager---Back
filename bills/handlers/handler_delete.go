package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/bills/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
	}

	bills := models.Bills{ID: id}

	deletedID, err := models.DeleteBills(bills)
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
