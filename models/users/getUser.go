package models

import (
	"github.com/Dnreikronos/budgetMannager---Back/db"
)

func Get(id int64) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM user WHERE id=$1`, user.ID)

	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role, &user.IsActive)

	return
}
