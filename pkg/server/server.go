package server

import (
	"adapptor-backend/pkg/api"
	"adapptor-backend/pkg/middleware"
	"fmt"
	"net/http"

	_ "adapptor-backend/docs" // import the swagger docs

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Adapptor Backend API
// @version 1.0
// @description This is a basic server for Adapptor Backend.
// @host localhost:8080
// @BasePath /
// @schemes http https

type Server struct {
	handler http.Handler
}

func New() *Server {
	router := mux.NewRouter()

	// Create CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Apply middlewares
	router.Use(middleware.Logging)

	// Register routes
	// note: granular control for auth on each route
	router.Handle("/welcome", middleware.Auth(api.HandleWelcome)).Methods("GET", "OPTIONS")
	router.Handle("/action", middleware.Auth(api.HandleAction)).Methods("POST", "OPTIONS")

	// Serve swagger files
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler())

	// Create handler chain with CORS
	handler := corsMiddleware.Handler(router)

	return &Server{
		handler: handler,
	}
}

func (s *Server) Start(port int) error {
	fmt.Printf("Starting server on port %d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.handler)
}
