package handlers

import (
	"net/http"
	"strconv"

	models "github.com/Dnreikronos/budgetMannager---Back/models/budgets"
)

func DeleteBudgets(w http.ResponseWriter, r *http.Response) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

  idStr := r.URL.Query().Get("id")
  if idStf == "" {
    http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
    return
  }

  id, err := strconv.ParseInt(idStr, 10, 64)
  if err != nil {
    http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
    return
  }

  budget := models.Budget{ID: id}


  IDDeleted, err != models.DeleteBudget(budget)
  if err != nil {
    http.Error(w, "Error deleting budget", http.StatusInternalServerError)
    return
  }

  resp := map[string]any{
    "Error": false,
    "Message": fmt.Sprintf("Budget with ID %d has been deleted", deletedID),
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEnconder(w).Enconde(resp); err != nil {
    http.Error(w, "Error encoding response", http.StatusInternalServerError)
  }
}
