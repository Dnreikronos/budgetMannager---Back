package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbUser     = os.Getenv("${DB_USER}")
	dbPassword = os.Getenv("${DB_PASSWORD}")
	dbName     = os.Getenv("budgetmannager")
	dbHost     = os.Getenv("${DB_HOST}")
	dbPort     = os.Getenv("5432")
)

func OpenConnection() (*gorm.DB, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }

	log.Println("Sucessfuly connected to the database!")
	return db, nil
}
