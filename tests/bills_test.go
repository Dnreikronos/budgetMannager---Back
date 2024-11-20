package handlers_test

import (
	"github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Register handlers
	r.PUT("/CreateBill", handlers.CreateBillsHandler)
	r.PUT("/Bill/:id", handlers.UpdateBillsHandler)
	r.DELETE("/Bill/:id", handlers.DeleteBillsHandler)
	r.GET("/Bill/:id", handlers.GetBillHandler)
	r.GET("/Bills", handlers.GetAllBillsHanddler)

	return r
}

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory;"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	if err := db.AutoMigrate(&models.Bills{}); err != nil {
		panic("Failed to migrate database")
	}
	return db
}
