package handlers

import (
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBudgetHandler(c *gin.Context) {
	var input models.Budget
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	newBudget := models.Budget{
		Value:    input.Value,
		Currency: input.Currency,
		Start:    input.Start,
		End:      input.End,
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
	}
	if err := db.Create(&newBudget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to create Budget"})
	}
	c.JSON(http.StatusCreated, gin.H{"Budget": newBudget})
}
