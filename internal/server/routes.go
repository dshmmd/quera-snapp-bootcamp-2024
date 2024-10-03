package server

import (
	"github.com/dshmmd/quera-snapp-bootcamp-2024/pkg/week1/q1"
	"github.com/dshmmd/quera-snapp-bootcamp-2024/pkg/week1/q2"
	"github.com/dshmmd/quera-snapp-bootcamp-2024/pkg/week1/q3"
	"github.com/dshmmd/quera-snapp-bootcamp-2024/pkg/week1/q5"
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/week1/q1", ResolutionHandler(w1q1.Solve))
	mux.HandleFunc("/week1/q2", ResolutionHandler(w1q2.Solve))
	mux.HandleFunc("/week1/q3", ResolutionHandler(w1q3.Solve))
	mux.HandleFunc("/week1/q5", ResolutionHandler(w1q5.Solve))
}
