package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"

	"palestra-go/config"
	"palestra-go/pkg/collaborator"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	dialInfo, err := mgo.ParseURL(config.MongoDBHost)
	dialInfo.Direct = true
	dialInfo.FailFast = true
	dialInfo.Database = config.MongoDBDatabase
	// dialInfo.Username = "admin"
	// dialInfo.Password = "admin"

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	r := mux.NewRouter()

	mPool := mgosession.NewPool(nil, session, config.MongoDBConnectionPool)
	defer mPool.Close()

	collaboratorRepo := collaborator.NewMongoRepository(mPool, config.MongoDBDatabase)
	collaboratorService := collaborator.NewService(collaboratorRepo)

	n := negroni.New(
		// negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	CreateRoutes(r, *n, collaboratorService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.APIPort),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
