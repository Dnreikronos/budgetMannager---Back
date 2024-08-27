package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func updateUser(user User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	sql := `UPDATE user
          set email = $2, name = $3, password = $3 , role = $5, is_active = $6
          where id = $1`

	_, err = conn.Exec(sql, user.ID, user.Email, user.Name, user.Password, user.Role, user.IsActive)
	return
}
