package collaborator

import (
	"palestra-go/pkg/entity"

	"gopkg.in/mgo.v2/bson"
)

// Reader methods that read Collaborators from storage
type Reader interface {
	Find(id bson.ObjectId) (*entity.Collaborator, error)
	Search(query string) ([]*entity.Collaborator, error)
	FindAll() ([]*entity.Collaborator, error)
}

// Writer methods that write Collaborators from storage
type Writer interface {
	Save(c *entity.Collaborator) (bson.ObjectId, error)
	Delete(id bson.ObjectId) error
}

// Repository aggregation of all interfaces comunicating with storage
type Repository interface {
	Reader
	Writer
}
