package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestCreateBillsHandler(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter(db)

	billInput := models.BillInput{
		Value:    1200,
		Category: "Electricity",
		Status:   "pending",
	}
	jsonData, _ := json.Marshal(billInput)

	req, _ := http.NewRequest(http.MethodPut, "/CreateBill", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code 201, got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if response["Bill"] == nil {
		t.Errorf("expected bill in response, got nil")
	}
}
