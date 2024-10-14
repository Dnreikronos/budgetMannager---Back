package main

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/configs"
	conn "github.com/Dnreikronos/budgetMannager---Back/db/connetion"
	"github.com/Dnreikronos/budgetMannager---Back/handlers"
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
    panic (err)
  }

 conn.RunMigrations(db)


	r := gin.Default()

	// corsMiddleware := cors.New(cors.Options{
	// 	AllowedOrigins: []string{
	// 		"http://localhost:5173",
	// 		"http://localhost:9000",
	// 	},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: true,
	// 	MaxAge:           300,
	// })
	//

	r.POST("/register", handlers.CreateUserHandler)
	r.POST("/login", handlers.LoginHandler)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
