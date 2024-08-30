package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func GetAll() (user []User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`Select * FROM user`)
	if err != nil {
		return
	}
	for rows.Next() {
		var users User
		err = rows.Scan(&users.ID, &users.Email, &users.Name, &users.Password, &users.Role, &users.IsActive)
		if err != nil {
			continue
		}
		user = append(user, users)
	}
	return
}
