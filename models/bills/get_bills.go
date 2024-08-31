package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func GetBills(bills Bills) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM bills WHERE id=$1`, bills.ID)

	err = row.Scan(&bills.ID, &bills.Value, &bills.UserID, &bills.BudgetID, &bills.Category, &bills.Status, &bills.CreatedAt, &bills.UpdatedAt)

	return
}
