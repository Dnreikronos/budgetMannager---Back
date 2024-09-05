package main

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/budgetMannager---Back/configs"
	"github.com/Dnreikronos/budgetMannager---Back/users/handlers"
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
	r.Get("/{id}", handlers.GetUser)
	r.Put("/updateUser", handlers.UpdateUser)
	r.Delete("/deleteUser", handlers.DeleteUser)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
