package tests

import (
	"github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


	func setupRouterBills() (*gin.Engine, *gorm.DB) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Bills{})

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	router.POST("/createBill", handlers.CreateBillsHandler)
	router.PUT("/Bill/id:", handlers.UpdateBillsHandler)
	router.DELETE("/Bill/id:", handlers.DeleteBillsHandler)
	router.GET("/Bill/id:", handlers.GetBillHandler)
	router.GET("/Bills", handlers.GetAllBillsHanddler)
	return router, db
}
