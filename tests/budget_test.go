package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	h "github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	r.POST("/budget", h.CreateBudgetHandler)
	r.PUT("/budget/:id", h.UpdateBudgetHandler)
	r.DELETE("/budget/:id", h.DeleteBudgetHandler)
	r.GET("/budget/:id", h.GetBudgetHandler)
	r.GET("/budgets", h.GetAllBudgetHandler)

	return r
}

func setupBudgetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
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

	req, _ := http.NewRequest(http.MethodPost, "/budget", bytes.NewBuffer(jsonData))
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

func TestUpdateBudgetsHandler(t *testing.T) {
	db := setupBudgetDB()
	router := setupBudgetRouter(db)

	testBudget := models.Budget{
		ID:       uuid.New(),
		Value:    1500,
		Currency: "BRL",
		Start:    time.Now(),
		End:      time.Now(),
	}
	if err := db.Create(&testBudget).Error; err != nil {
		t.Fatalf("Failed to create test budget: %v", err)
	}

	updatedInput := models.BudgetInput{
		Value:    1600,
		Currency: "USD",
		Start:    time.Now(),
		End:      time.Now(),
	}
	jsonData, _ := json.Marshal(updatedInput)

	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/budget/%s", testBudget.ID.String()), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	responseBudget := response["Budget"].(map[string]interface{})
	if int64(responseBudget["value"].(float64)) != updatedInput.Value {
		t.Errorf("Expected updated value %v, got %v", updatedInput.Value, responseBudget["value"])
	}
}

func TestDeleteBudgetsHandler(t *testing.T) {
	db := setupBudgetDB()
	router := setupBudgetRouter(db)

	testBudget := models.Budget{
		ID:       uuid.New(),
		Value:    1500,
		Currency: "BRL",
		Start:    time.Now(),
		End:      time.Now(),
	}
	if err := db.Create(&testBudget).Error; err != nil {
		t.Fatalf("Failed to create test budget: %v", err)
	}

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/budget/%s", testBudget.ID.String()), nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["Status"] != "Budget deleted with sucess!" { // Note the typo in "sucess"
		t.Errorf("Expected success message, got %v", response["Status"])
	}

	var count int64
	db.Model(&models.Budget{}).Where("id = ?", testBudget.ID).Count(&count)
	if count != 0 {
		t.Errorf("Expected budget to be deleted, but it still exists")
	}
}

func TestGetBudgetsHandler(t *testing.T) {
	db := setupBudgetDB()
	router := setupBudgetRouter(db)

	testBudget := models.Budget{
		ID:       uuid.New(),
		Value:    1500,
		Currency: "BRL",
		Start:    time.Now(),
		End:      time.Now(),
	}
	if err := db.Create(&testBudget).Error; err != nil {
		t.Fatalf("Failed to create test budget: %v", err)
	}

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/budget/%s", testBudget.ID.String()), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	responseBudget := response["Budget"].(map[string]interface{})
	if responseBudget["currency"] != testBudget.Currency {
		t.Errorf("Expected currency %v, got %v", testBudget.Currency, responseBudget["currency"])
	}
}

func TestGetAllBudgetsHandler(t *testing.T) {
	db := setupBudgetDB()
	router := setupBudgetRouter(db)

	testBudgets := []models.Budget{
		{ID: uuid.New(), Value: 100, Currency: "USD"},
		{ID: uuid.New(), Value: 200, Currency: "BRL"},
	}

	for _, budget := range testBudgets {
		if err := db.Create(&budget).Error; err != nil {
			t.Fatalf("Failed to create test budget: %v", err)
		}
	}

	req, _ := http.NewRequest(http.MethodGet, "/budgets", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	budgets, ok := response["Budgets"].([]interface{})
	if !ok {
		t.Fatalf("Expected array of budgets, got %v", response["Budgets"])
	}

	if len(budgets) != len(testBudgets) {
		t.Errorf("Expected %d budgets, got %d", len(testBudgets), len(budgets))
	}
}
