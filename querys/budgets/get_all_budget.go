package querys

import (
	"github.com/Dnreikronos/budgetMannager---Back/db"
	"github.com/Dnreikronos/budgetMannager---Back/models"
)

func GetAll() (budget []models.Budget, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM budget`)
	if err != nil {
		return
	}
	for rows.Next() {
		var budgets models.Budget
		err = rows.Scan(&budgets.ID, &budgets.Value, &budgets.UserID, &budgets.Currency, &budgets.Validity)
		if err != nil {
			continue
		}
		budget = append(budget, budgets)
	}

	return
}
