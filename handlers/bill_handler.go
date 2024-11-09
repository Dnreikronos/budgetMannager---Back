package handlers

import (
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBillsHandler(c *gin.Context) {
	var input models.BillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	newBill := models.Bills{
		Value:    input.Value,
		Category: input.Category,
		Status:   input.Status,
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
	}

	if err := db.Create(&newBill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to create Bills"})
	}
	c.JSON(http.StatusCreated, gin.H{"Bill": newBill})
}
