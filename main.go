package main

import (
	"fmt"
	"log"

	"github.com/juju/mgosession"

	"palestra-go/config"
	"palestra-go/pkg/collaborator"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial(config.MongoDBHost)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, config.MongoDBConnectionPool)
	defer mPool.Close()

	collaboratorRepo := collaborator.NewMongoRepository(mPool, config.MongoDBDatabase)
	// id, err := collaboratorRepo.Save(&entity.Collaborator{
	// 	ID:        bson.NewObjectId(),
	// 	Name:      "José da Silva",
	// 	Email:     "josesilva@gmail.com",
	// 	Social:    entity.CollabSocial{Homepage: "page", Github: "github", Twitter: "twitter"},
	// 	CreatedAt: time.Now(),
	// })
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println(id)

	result, _ := collaboratorRepo.FindAll()

	for _, el := range result {
		fmt.Println(el)
	}
}
