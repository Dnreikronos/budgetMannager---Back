package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func GetAll() (bills []Bills, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM bills`)
	if err != nil {
		return
	}
	for rows.Next() {
		var bill Bills
		err = rows.Scan(&bill.ID, &bill.Value, &bill.UserID, &bill.BudgetID, &bill.Category, &bill.Status, &bill.CreatedAt, &bill.UpdatedAt)
		if err != nil {
			continue
		}
		bills = append(bills, bill)
	}
	return
}
