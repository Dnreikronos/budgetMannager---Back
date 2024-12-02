package tests

import (
	"log"

	h "github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupBudgetRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.POST("/CreateBudget", h.CreateBudgetHandler)
	r.PUT("/Bill/:id", h.UpdateBudgetHandler)
	r.DELETE("/Bill/:id", h.DeleteBudgetHandler)
	r.GET("/Bill/:id", h.GetBudgetHandler)
	r.GET("/Budgets", h.GetAllBudgetHandler)

	return r
}

func setupBudgetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory;"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	err = db.Migrator().DropTable(&models.Budget{})
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
	if err := db.AutoMigrate(&models.Budget{}); err != nil {
		panic("Failed to migrate database")
	}
	return db
}
