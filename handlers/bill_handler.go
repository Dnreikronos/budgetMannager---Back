package handlers

import (
	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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

func UpdateBillsHandler(c *gin.Context) {
	var input models.BillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": err.Error()})
		return
	}

	billID := c.Param("id")
	if billID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Bill ID is required"})
		return
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database connection error"})
		return
	}

	var bill models.Bills
	if err := db.First(&bill, "id = ?", billID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"Status": "Bill not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": "Failed to fetch Bill"})
		}
		return
	}

	bill.Value = input.Value
	bill.Category = input.Category
	bill.Status = input.Status
	if err := db.Save(&bill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Failed to update Bill"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Bill": bill})
}

func DeleteBillsHandler(c *gin.Context) {
	billID := c.Param("ID")
	if billID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Bill ID is required"})
		return
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database connection error"})
		return
	}
	var bill models.Bills
	if err := db.First(&bill, "id = ?", billID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"Status": "Bill not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": "Failed to fetch Bill"})
		}
		return
	}

	if err := db.Delete(&bill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Failed to delete Bill"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Bill deleted with sucess!"})
}

