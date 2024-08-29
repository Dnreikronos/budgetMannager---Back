package models

import "github.com/Dnreikronos/budgetMannager---Back/db"

func getAll() (users []User) {
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
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role, &user.IsActive)
		if err != nil {
			continue
		}
		users = append(users, user)
	}
	return
}
