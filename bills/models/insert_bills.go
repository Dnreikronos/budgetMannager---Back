package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func CreateBills(bills Bills) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO bills (id, value, user_id, budget_id, category, status, created_at, updated_at)`

	err = conn.QueryRow(sql, bills.ID, bills.Value, bills.UserID, bills.BudgetID, bills.Category, bills.Status, bills.CreatedAt, bills.UpdatedAt).Scan(&id)

	return
}
