package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CollabLoginData A temporary structure to create a simple login
type CollabLoginData struct {
	Username string        `json:"username"`
	Pwd      string        `json:"pwd"`
	CollabID bson.ObjectId `json:"id"`
}

// Collaborator User data
type Collaborator struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Social    CollabSocial  `json:"social"`
	CreatedAt time.Time     `json:"createdAt"`
}

// CollabSocial An embedded object inside Collaborator
type CollabSocial struct {
	Homepage string `json:"homepage"`
	Github   string `json:"github"`
	Twitter  string `json:"twitter"`
}
