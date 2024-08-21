package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/user", makeHttpHandleFunc(s.handleUser))
	router.HandleFunc("/budget", makeHttpHandleFunc(s.handleBudget))
	router.HandleFunc("/bills", makeHttpHandleFunc(s.handleBills))

	log.Println("Json API running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)

}

type ApiError struct {
	Error string
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUser(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}

	return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleBudget(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		s.handleGetBudget(w, r)
	}
	if r.Method == "POST" {
		s.handleCreateBudget(w, r)
	}
	if r.Method == "DELETE" {
		s.handleDeleteBudget(w, r)
	}
	if r.Method == "UPDATE" {
		s.handleUpdateBudget(w, r)
	}
	return nil
}

func (s *APIServer) handleGetBudget(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateBudget(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateBudget(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteBudget(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleBills(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		s.handleGetBills(w, r)
	}
	if r.Method == "POST" {
		s.handleCreateBills(w, r)
	}
	if r.Method == "DELETE" {
		s.handleDeleteBills(w, r)
	}
	if r.Method == "UPDATE" {
		s.handleUpdateBills(w, r)
	}
	return nil
}

func (s *APIServer) handleGetBills(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateBills(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteBills(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateBills(w http.ResponseWriter, r *http.Request) error {
	return nil
}
