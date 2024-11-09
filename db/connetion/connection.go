package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost         = os.Getenv("DB_HOST")
	dbPort         = os.Getenv("DB_PORT")
	dbUser         = os.Getenv("DB_USER")
	dbPassword     = os.Getenv("DB_PASSWORD")
	dbName         = os.Getenv("POSTGRES_DB")
	dbTimeZone     = os.Getenv("POSTGRES_TIME_ZONE")

)

func OpenConnection() (*gorm.DB, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
		dbTimeZone,)

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
