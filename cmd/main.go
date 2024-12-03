package main

import (
	"fmt"
	"github.com/Dnreikronos/budgetMannager---Back/configs"
	conn "github.com/Dnreikronos/budgetMannager---Back/db/connetion"
	h "github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	err = configs.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	db, err := conn.OpenConnection()
	if err != nil {
		panic(err)
	}

	conn.RunMigrations(db)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	corsOrigin := os.Getenv("CORS")

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{corsOrigin},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	//User
	r.POST("/register", h.CreateUserHandler)
	r.POST("/login", h.LoginHandler)
	authorized := r.Group("/", h.AuthMiddleware())
	{
		authorized.GET("/profile", h.ProfileHandler)
	}

	//Bills
	r.POST("/CreateBill", h.CreateBillsHandler)
	r.PUT("/Bill/:id", h.UpdateBillsHandler)
	r.DELETE("/Bill/:id", h.DeleteBillsHandler)
	r.GET("/Bill/:id", h.GetBillHandler)
	r.GET("/Bills", h.GetAllBillsHanddler)

	//Budgets
	r.POST("/CreateBudget", h.CreateBudgetHandler)
	r.PUT("/Budget/:id", h.UpdateBudgetHandler)
	r.DELETE("/Budget/:id", h.DeleteBudgetHandler)
	r.GET("/Budget/:id", h.GetBudgetHandler)
	r.GET("/Budgets", h.GetAllBudgetHandler)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
