package querys

import (
	"github.com/Dnreikronos/budgetMannager---Back/db"
	"github.com/Dnreikronos/budgetMannager---Back/models"
)

func Get(id int64) (user models.User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM user WHERE id=$1`, user.ID)

	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role, &user.IsActive)

	return
}
