package db 

import (
	m "github.com/Dnreikronos/budgetMannager---Back/models"
	"gorm.io/gorm"
)


func RunMigrations(db *gorm.DB) {
  createTables(db)
}

func createTables(db *gorm.DB) {
  db.AutoMigrate(&m.User{})
  db.AutoMigrate(&m.Bills{})
  db.AutoMigrate(&m.Budget{})
}
