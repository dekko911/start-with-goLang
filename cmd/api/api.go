package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dekko911/start-with-goLang/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	db      *sql.DB
}

// Creating a new Instance of the API Server.
func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

// Running & Listening the server. The 's' stands for server.
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)

	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on port " + s.address)

	return http.ListenAndServe(s.address, router)
}
