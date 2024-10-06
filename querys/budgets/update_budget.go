package querys

import (
	"github.com/Dnreikronos/budgetMannager---Back/db"
	"github.com/Dnreikronos/budgetMannager---Back/models"
)


func UpdateBudget(budget models.Budget) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `UPDATE budget
					set value = $2, userid = $3, currency = $4, validity = $5
					where id = $1`

	_, err = conn.Exec(sql, budget.ID, budget.Validity, budget.Currency, budget.Validity)

	return
}
