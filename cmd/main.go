package main

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/configs"
	"github.com/Dnreikronos/budgetMannager---Back/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://localhost:9000",
			"*",
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
	r.Get("/GetUser{id}", handlers.GetUser)
	r.Put("/UpdateUser", handlers.UpdateUser)
	r.Delete("/deleteUser", handlers.DeleteUser)

	r.Post("/CreateBudget", handlers.CreateBudget)
	r.Get("/GetAllBudget", handlers.GetAllBudget)
	r.Get("/GetBudget{id}", handlers.GetBudget)
	r.Put("/UpdateBudget", handlers.UpdateBudget)
	r.Delete("/DeleteBudgets", handlers.DeleteBudget)

	r.Post("/CreateBills", handlers.CreateBills)
	r.Get("/GetAllBills", handlers.GetAllBills)
	r.Get("/GetBills{id}", handlers.GetBills)
	r.Put("/UpdateBills", handlers.UpdateBills)
	r.Delete("/DeleteBills", handlers.DeleteBills)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
