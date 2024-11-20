package handlers_test

import (
	"github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/gin-gonic/gin"
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


