package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
	r.PUT("/Budget/:id", h.UpdateBudgetHandler)
	r.DELETE("/Budget/:id", h.DeleteBudgetHandler)
	r.GET("/Budget/:id", h.GetBudgetHandler)
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

func TestCreateBudgetHandler(t *testing.T) {
	db := setupBudgetDB()
	router := setupBudgetRouter(db)

	budgetInput := models.BudgetInput{
		Value:    1500,
		Currency: "USD",
		Start:    time.Now(),
		End:      time.Now(),
	}
	jsonData, _ := json.Marshal(budgetInput)

	req, _ := http.NewRequest(http.MethodPost, "/CreateBudget", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, but got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["Budget"] == nil {
		t.Errorf("Expected budget in response, got nil")
	}
}
