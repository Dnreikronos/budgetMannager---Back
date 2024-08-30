package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func InsertUser(user User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO user (id, email, name, password, role, is_active) VALUES ($1, $2, $3, $4, $5, $6) RETURN id`

	err = conn.QueryRow(sql, user.ID, user.Email, user.Name, user.Password, user.Role, user.IsActive).Scan(&id)

	return
}
