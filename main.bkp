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
