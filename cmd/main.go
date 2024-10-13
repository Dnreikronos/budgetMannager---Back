package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/configs"
	"github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := configs.Load(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	serverPort := configs.GetServerPort()
	log.Printf("Server port: %s", serverPort)

	dbConfig := configs.GetDB()
	log.Printf("Config DB: %s", dbConfig)

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
