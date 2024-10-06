package querys

import (
	"github.com/Dnreikronos/budgetMannager---Back/db"
	"github.com/Dnreikronos/budgetMannager---Back/models"
)

func DeleteBills(bills models.Bills) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `DELETE FROM bills where id = $1`

	result, err := conn.Exec(sql, bills.ID)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected > 1 {
		id = bills.ID
	} else {
		id = 0
	}

	return
}
