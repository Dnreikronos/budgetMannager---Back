package main

import "net/http"

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	return nil
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
