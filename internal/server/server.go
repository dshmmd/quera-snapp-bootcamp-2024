package server

import (
	"context"
	"fmt"
	"github.com/dshmmd/quera-snapp-bootcamp-2024/internal/resolver"
	"log"
	"net/http"
	"time"

	"github.com/dshmmd/quera-snapp-bootcamp-2024/config"
)

// Server struct containing the custom http.Server
type Server struct {
	httpServer *http.Server
}

// NewServer creates a new instance of Server with a custom http.Server
func NewServer(cfg *config.Config) (*Server, error) {
	// Create a new ServeMux to handle different routes
	mux := http.NewServeMux()

	// Register routes with their handlers
	mux.HandleFunc("/healthz", healthCheckHandler)
	registerRoutes(mux)

	// Create a custom http.Server with the given address and handler (mux)
	s := Server{
		httpServer: &http.Server{
			Addr:         cfg.BindAddress,
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
	}

	return &s, nil
}

// Serve sets up routes on the Server's custom http.Server and starts it
func (s *Server) Serve() error {
	log.Printf("Starting HTTP Server at %s\n", s.httpServer.Addr)
	// Start the server
	return s.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the server with a timeout
func (s *Server) Stop(ctx context.Context) error {
	if s.httpServer != nil {
		return s.httpServer.Shutdown(ctx)
	}
	return nil
}

// healthCheckHandler handles requests to the /healthz path
func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Health check: OK")
}

func ResolutionHandler(slv resolver.Resolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		answer, err := slv(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(answer))
	}
}
