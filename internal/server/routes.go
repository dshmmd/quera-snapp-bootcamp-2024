package server

import (
	"net/http"

	"github.com/dshmmd/quera-snapp-bootcamp-2024/pkg/week1/q1"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/week1/q1", ResolutionHandler(w1q1.Solve))
}
