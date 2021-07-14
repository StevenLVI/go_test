// internal/server/server.go
package server

import (
	v1 "go_test/internal/server/v1"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()

	// API routes version 1.
	r.Mount("/api/v1", v1.New())

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Close server resources.
func (serv *Server) Close() error {
	// TODO: add resource closure.
	return nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
