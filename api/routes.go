package main

import (
	"palestra-go/api/handlers/collaborators"
	"palestra-go/pkg/collaborator"

	"github.com/go-chi/chi"
)

// CreateRoutes .
func CreateRoutes(r chi.Router, service collaborator.UseCase) {
	r.Route("/collaborators", func(r chi.Router) {
		r.Use(collaborators.ServiceContext(&service))

		r.Get("/", collaborators.Get)  // GET /collaborators
		r.Post("/", collaborators.Add) // POST /collaborators

		r.Route("/{id}", func(r chi.Router) {
			r.Use(collaborators.Context)

			r.Get("/", collaborators.GetOne)    // GET /collaborators/{id}
			r.Put("/", collaborators.Put)       // PUT /collaborators/{id}
			r.Delete("/", collaborators.Delete) // DELETE /collaborators/{id}
		})
	})
}
