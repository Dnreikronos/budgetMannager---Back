package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func deleteUser(user User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `DELETE * FROM user where id = $1`

	_, err = conn.Exec(sql)

	return
}
