package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func DeleteBudget(budget Budget) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `DELETE FROM budget WHERE id = $1`

	result, err := conn.Exec(sql, budget.ID)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected > 0 {
		id = budget.ID
	} else {
		id = 0
	}

	return
}
