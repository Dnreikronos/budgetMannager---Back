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

func UpdateBudgetHandler(c *gin.Context) {
	var input models.BudgetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": err.Error()})
		return
	}

	budgetID := c.Param("id")
	if budgetID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Budget ID is required"})
		return
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database connection error"})
		return
	}

	var budget models.Budget
	if err := db.First(&budget, "id = ?", budgetID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"Status": "Budget not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": "Failed to fetch Budget"})
		}
		return
	}

	budget.Value = input.Value
	budget.Currency = input.Currency
	budget.Start = input.Start
	budget.End = input.End
	if err := db.Save(&budget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Failed to update Budget"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Budget": budget})
}

func DeleteBudgetHandler(c *gin.Context) {
	budgetID := c.Param("id")
	if budgetID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Budget ID is required"})
		return
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database connection error"})
		return
	}

	var budget models.Budget
	if err := db.First(&budget, "id = ?", budgetID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"Status": "Budget not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": "Failed to fetch Budget"})
		}
		return
	}

	if err := db.Delete(&budget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Failed to delete Budget"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Budget deleted with sucess!"})
}
