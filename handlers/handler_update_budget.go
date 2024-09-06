package handlers

import (
	"fmt"
	"net/http"
)

func UpdateBudget(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Register with ID %s has been updated", id)
}
