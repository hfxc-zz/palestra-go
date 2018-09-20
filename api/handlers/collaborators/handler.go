package collaborators

import (
	"encoding/json"
	"log"
	"net/http"
	"palestra-go/pkg/entity"

	"github.com/go-chi/chi"

	"gopkg.in/mgo.v2/bson"
)

// Get list all collaborators
func Get(w http.ResponseWriter, r *http.Request) {
	service, ok := Service(r)
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

// Add create a new collaborator
func Add(w http.ResponseWriter, r *http.Request) {
	service, ok := Service(r)
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

// GetOne return one collaborator
func GetOne(w http.ResponseWriter, r *http.Request) {
	collaborator, ok := Collaborator(r)
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

// Delete remove one collaborator
func Delete(w http.ResponseWriter, r *http.Request) {
	service, ok := Service(r)
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

// Put edit one collaborator
func Put(w http.ResponseWriter, r *http.Request) {
	service, ok := Service(r)
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	collaborator, ok := Collaborator(r)
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var requestEntity *entity.Collaborator

	if err := json.NewDecoder(r.Body).Decode(&requestEntity); err != nil {
		log.Println(err.Error())
		return
	}

	if !requestEntity.Valid() {
		log.Println("err: invalid entity passed as request body")
		return
	}

	collaborator.CompareAndUpdate(requestEntity)

	if err := (*service).Update(collaborator); err != nil {
		http.Error(w, http.StatusText(http.StatusNotModified), http.StatusNotModified)
		return
	}

	responseEntity, err := json.Marshal(collaborator)
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseEntity)
}

// Patch edit one collaborator
func Patch(w http.ResponseWriter, r *http.Request) {
	service, ok := Service(r)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	collaborator, ok := Collaborator(r)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	var requestEntity *entity.Collaborator

	if err := json.NewDecoder(r.Body).Decode(&requestEntity); err != nil {
		log.Println(err.Error())
		return
	}

	collaborator.CompareAndUpdate(requestEntity)

	if err := (*service).Update(collaborator); err != nil {
		http.Error(w, http.StatusText(http.StatusNotModified), http.StatusNotModified)
		return
	}

	cJSON, err := json.Marshal(collaborator)
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cJSON)
}
