package collaborators

import (
	"context"
	"net/http"
	"palestra-go/pkg/collaborator"
	"palestra-go/pkg/entity"

	"github.com/go-chi/chi"
	"gopkg.in/mgo.v2/bson"
)

type contextKey string

// ServiceContext adds the service to router context
func ServiceContext(service *collaborator.UseCase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), contextKey("service"), service)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// EntityContext tries to load the collaborator specified in the URLParam and add it to router context,
// if its not possible an error is written in the response
func EntityContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxKey := contextKey("service")
		service, ok := ctx.Value(ctxKey).(*collaborator.UseCase)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		collaborator, err := (*service).Find(bson.ObjectIdHex(id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx = context.WithValue(r.Context(), contextKey("collaborator"), collaborator)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Service returns the service in the router context
func Service(r *http.Request) (*collaborator.UseCase, bool) {
	ctx := r.Context()
	ctxKey := contextKey("service")
	service, ok := ctx.Value(ctxKey).(*collaborator.UseCase)
	return service, ok
}

// Collaborator returns the collaborator in the router context
func Collaborator(r *http.Request) (*entity.Collaborator, bool) {
	ctx := r.Context()
	ctxKey := contextKey("collaborator")
	collaborator, ok := ctx.Value(ctxKey).(*entity.Collaborator)
	return collaborator, ok
}
