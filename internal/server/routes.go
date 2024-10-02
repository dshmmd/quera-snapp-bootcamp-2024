package server

import (
	"github.com/dshmmd/quera-snapp-bootcamp-2024/pkg/week1/q1"
	"github.com/dshmmd/quera-snapp-bootcamp-2024/pkg/week1/q2"
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/week1/q1", ResolutionHandler(w1q1.Solve))
	mux.HandleFunc("/week1/q2", ResolutionHandler(w1q2.Solve))
}
