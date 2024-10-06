package querys

import (
	"github.com/Dnreikronos/budgetMannager---Back/db"
	"github.com/Dnreikronos/budgetMannager---Back/models"
)


func GetBills(id int64) (bills models.Bills, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM bills WHERE id=$1`, bills.ID)

	err = row.Scan(&bills.ID, &bills.Value, &bills.UserID, &bills.BudgetID, &bills.Category, &bills.Status, &bills.CreatedAt, &bills.UpdatedAt)

	return
}
