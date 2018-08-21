package collaborators

import (
	"context"
	"encoding/json"
	"log"
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

// Context tries to load the collaborator specified in the URLParam if not possible an error is shown
func Context(next http.Handler) http.Handler {
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

// Get .
func Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctxKey := contextKey("service")
	service, ok := ctx.Value(ctxKey).(*collaborator.UseCase)
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := (*service).FindAll()
	if err != nil && err != entity.ErrNotFound {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if data == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// Add .
func Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctxKey := contextKey("service")
	service, ok := ctx.Value(ctxKey).(*collaborator.UseCase)
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var c *entity.Collaborator
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("errorMessage"))
		return
	}

	c.ID, err = (*service).Create(c)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("errorMessage"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("errorMessage"))
		return
	}
}

// GetOne .
func GetOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctxKey := contextKey("collaborator")
	collaborator, ok := ctx.Value(ctxKey).(*entity.Collaborator)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(collaborator); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error converting collaborator"))
	}
}

// Delete .
func Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctxKey := contextKey("service")
	service, ok := ctx.Value(ctxKey).(*collaborator.UseCase)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	id := chi.URLParam(r, "id")

	if err := (*service).Delete(bson.ObjectIdHex(id)); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Put .
func Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
