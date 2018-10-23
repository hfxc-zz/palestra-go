package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/juju/mgosession"

	"palestra-go/config"
	"palestra-go/pkg/collaborator"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	// Configure database connection info
	dialInfo, err := mgo.ParseURL(config.MongoDBHost)
	dialInfo.Direct = true
	dialInfo.FailFast = true
	dialInfo.Database = config.MongoDBDatabase

	// Connect to database
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, config.MongoDBConnectionPool)
	defer mPool.Close()

	// Initialize repository and service
	collaboratorRepo := collaborator.NewMongoRepository(mPool, config.MongoDBDatabase)
	collaboratorService := collaborator.NewService(collaboratorRepo)

	// Initialize and configure router
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	CreateRoutes(r, collaboratorService)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.APIPort),
		Handler:      r,
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
