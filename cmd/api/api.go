package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dekko911/start-with-goLang/service/product"
	"github.com/dekko911/start-with-goLang/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	db      *sql.DB
}

// Creating a new Instance of the API Server. alias construct.
func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

// This routes for api.
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// user routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// product routes
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on port", s.address)

	return http.ListenAndServe(s.address, router)
}
