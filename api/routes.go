package main

import (
	handler "palestra-go/api/handlers/collaborator"
	"palestra-go/pkg/collaborator"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// CreateRoutes .
func CreateRoutes(r *mux.Router, n negroni.Negroni, service collaborator.UseCase) {
	r.Handle("/collaborator", n.With(
		negroni.Wrap(handler.Root(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/collaborator/{id}", n.With(
		negroni.Wrap(handler.Get(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/collaborator", n.With(
		negroni.Wrap(handler.Add(service)),
	)).Methods("POST", "OPTIONS")
}
