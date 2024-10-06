package querys

import (
	"github.com/Dnreikronos/budgetMannager---Back/db"
	"github.com/Dnreikronos/budgetMannager---Back/models"
)

func DeleteUser(user models.User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `DELETE FROM user WHERE id = $1`

	result, err := conn.Exec(sql, user.ID)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected > 0 {
		id = user.ID
	} else {
		id = 0
	}

	return
}
