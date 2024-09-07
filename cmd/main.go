package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/configs"
	"github.com/Dnreikronos/budgetMannager---Back/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	if err := configs.Load(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	serverPort := configs.GetServerPort()
	log.Printf("Server port: %s", serverPort)

	dbConfig := configs.GetDB()
	log.Printf("Config DB: %s", dbConfig)

	r := chi.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://localhost:9000",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsMiddleware.Handler)

	r.Post("/createUser", handlers.CreateUser)
	r.Get("/getAllUsers", handlers.GetAllUsers)
	r.Get("/getUser/{id}", handlers.GetUser)
	r.Put("/updateUser", handlers.UpdateUser)
	r.Delete("/deleteUser", handlers.DeleteUser)

	r.Post("/createBudget", handlers.CreateBudget)
	r.Get("/getAllBudgets", handlers.GetAllBudget)
	r.Get("/getBudget/{id}", handlers.GetBudget)
	r.Put("/updateBudget", handlers.UpdateBudget)
	r.Delete("/deleteBudgets", handlers.DeleteBudget)

	r.Post("/createBills", handlers.CreateBills)
	r.Get("/getAllBills", handlers.GetAllBills)
	r.Get("/getBills/{id}", handlers.GetBills)
	r.Put("/updateBills", handlers.UpdateBills)
	r.Delete("/deleteBills", handlers.DeleteBills)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
