package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	
	"github.com/Dnreikronos/budgetMannager---Back/configs"
	conn "github.com/Dnreikronos/budgetMannager---Back/db/connetion"
	"github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := conn.OpenConnection()
	if err != nil {
		panic(err)
	}

	conn.RunMigrations(db)

	r := gin.Default()

	corsOrigin := os.Getenv("CORS")

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{corsOrigin},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	r.POST("/register", handlers.CreateUserHandler)
	r.POST("/login", handlers.LoginHandler)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
