package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	err = db.Migrator().DropTable(&models.Bills{})
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
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

func TestUpdateBillsHandler(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter(db)

	testBill := models.Bills{
		ID:       uuid.New(),
		Value:    1500,
		Category: "Rent",
		Status:   "unpaid",
	}
	if err := db.Create(&testBill).Error; err != nil {
		t.Fatalf("failed to create test bill: %v", err)
	}

	updatedInput := models.BillInput{
		Value:    1600,
		Category: "Updated Rent",
		Status:   "paid",
	}
	jsonData, _ := json.Marshal(updatedInput)

	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/Bill/%s", testBill.ID.String()), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code 200, got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	responseBill := response["Bill"].(map[string]interface{})
	if responseBill["value"] != float64(updatedInput.Value) {
		t.Errorf("expected updated value %v, got %v", updatedInput.Value, responseBill["value"])
	}
}

func TestDeleteBillsHandler(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter(db)

	testBill := models.Bills{
		ID:       uuid.New(),
		Value:    500,
		Category: "Groceries",
		Status:   "unpaid",
	}
	if err := db.Create(&testBill).Error; err != nil {
		t.Fatalf("failed to create test bill: %v", err)
	}

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/Bill/%s", testBill.ID.String()), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code 200, got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if response["Status"] != "Bill deleted with sucess!" {
		t.Errorf("expected success message, got %v", response["Status"])
	}

	var count int64
	db.Model(&models.Bills{}).Where("id = ?", testBill.ID).Count(&count)
	if count != 0 {
		t.Errorf("expected bill to be deleted, but it still exists")
	}
}

func TestGetBillHandler(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter(db)

	testBill := models.Bills{
		ID:       uuid.New(),
		Value:    2000,
		Category: "Insurance",
		Status:   "paid",
	}
	if err := db.Create(&testBill).Error; err != nil {
		t.Fatalf("failed to create test bill: %v", err)
	}

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/Bill/%s", testBill.ID.String()), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code 200, got %v", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	responseBill := response["Bill"].(map[string]interface{})
	if responseBill["category"] != testBill.Category {
		t.Errorf("expected category %v, got %v", testBill.Category, responseBill["category"])
	}
}

func TestGetAllBillsHandler(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter(db)

	testBills := []models.Bills{
		{ID: uuid.New(), Value: 100, Category: "Groceries", Status: "paid"},
		{ID: uuid.New(), Value: 200, Category: "Utilities", Status: "unpaid"},
	}

	var createdBillIDs []string
	for _, bill := range testBills {
		if err := db.Create(&bill).Error; err != nil {
			t.Fatalf("failed to create test bill: %v", err)
		}
		createdBillIDs = append(createdBillIDs, bill.ID.String())
	}

	req, _ := http.NewRequest(http.MethodGet, "/Bills", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code 200, got %d", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	bills, ok := response["Bills"].([]interface{})
	if !ok {
		t.Fatalf("expected array of bills, got %v", response["Bills"])
	}

	if len(bills) != len(testBills) {
		t.Errorf("expected %d bills, got %d", len(testBills), len(bills))
	}

	for i, bill := range bills {
		billData := bill.(map[string]interface{})
		billID := billData["id"].(string)

		found := false
		for _, createdID := range createdBillIDs {
			if billID == createdID {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("unexpected bill ID %s", billID)
		}

		fmt.Printf("Expected value type: %T, Actual value type: %T\n", testBills[i].Value, billData["value"])

		if billData["value"] != float64(testBills[i].Value) {
			t.Errorf("expected bill value %d, got %v", testBills[i].Value, billData["value"])
		}
		if billData["category"] != testBills[i].Category {
			t.Errorf("expected bill category %s, got %s", testBills[i].Category, billData["category"])
		}
		if billData["status"] != testBills[i].Status {
			t.Errorf("expected bill status %s, got %s", testBills[i].Status, billData["status"])
		}
	}
}
