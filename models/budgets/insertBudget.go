package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func InsertBudget(budget Budget) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO budget (id, value, userid, currency, validity) VALUES ($1, $2, $3, $4, $5) RETURN id`

	err = conn.QueryRow(sql, budget.ID, budget.Value, budget.UserID, budget.Currency, budget.Validity).Scan(&id)

	return
}
