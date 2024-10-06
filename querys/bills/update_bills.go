package querys

import (
	"github.com/Dnreikronos/budgetMannager---Back/db"
	"github.com/Dnreikronos/budgetMannager---Back/models"
)

func UpdateBills(bills models.Bills) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `UPDATE bills
					set value = $2, user_id = $3, budget_id = $4, category = $5, status = $6, created_at = $ 7, updated_at = $ 8
					where id = $1`

	_, err = conn.Exec(sql, bills.ID, bills.Value, bills.UserID, bills.BudgetID, bills.Category, bills.Status, bills.CreatedAt, bills.UpdatedAt)

	return

}
