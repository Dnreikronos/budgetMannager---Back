package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func Get(int int64) (budget Budget, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM budget WHERE id=$1`, budget.ID)

	err = row.Scan(&budget.ID, &budget.Value, &budget.UserID, &budget.Currency, &budget.Validity)

	return
}
