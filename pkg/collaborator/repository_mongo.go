package collaborator

import (
	"palestra-go/pkg/entity"

	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

// MongoRepository .
type MongoRepository struct {
	pool *mgosession.Pool
	db   string
}

// NewMongoRepository .
func NewMongoRepository(p *mgosession.Pool, db string) *MongoRepository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}

// Find finds a Collaborator by id
func (r *MongoRepository) Find(id bson.ObjectId) (*entity.Collaborator, error) {
	result := entity.Collaborator{}
	session := r.pool.Session(nil)
	collection := session.DB(r.db).C("collaborator")
	err := collection.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

// FindAll list all Collaborators
func (r *MongoRepository) FindAll() ([]*entity.Collaborator, error) {
	result := []*entity.Collaborator{}
	session := r.pool.Session(nil)
	collection := session.DB(r.db).C("collaborator")
	err := collection.Find(nil).Sort("createdAt").All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

// Search find a Collaborator by his name
func (r *MongoRepository) Search(query string) ([]*entity.Collaborator, error) {
	result := []*entity.Collaborator{}
	session := r.pool.Session(nil)
	collection := session.DB(r.db).C("collaborator")
	err := collection.Find(bson.M{"name": &bson.RegEx{Pattern: query, Options: "i"}}).Limit(10).Sort("name").All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

// Create save a new Collaborator in database
func (r *MongoRepository) Create(c *entity.Collaborator) (bson.ObjectId, error) {
	session := r.pool.Session(nil)
	collection := session.DB(r.db).C("collaborator")
	err := collection.Insert(c)
	if err != nil {
		return bson.ObjectId(0), err
	}
	return c.ID, nil
}

// Update updates an existent Collaborator in database
func (r *MongoRepository) Update(c *entity.Collaborator) error {
	session := r.pool.Session(nil)
	collection := session.DB(r.db).C("collaborator")
	err := collection.Update(bson.M{"_id": c.ID}, c)
	switch err {
	case nil:
		return nil
	case mgo.ErrNotFound:
		return entity.ErrNotFound
	default:
		return err
	}
}

// Delete delete a Collaborator with the id passed as parameter
func (r *MongoRepository) Delete(id bson.ObjectId) error {
	session := r.pool.Session(nil)
	collection := session.DB(r.db).C("collaborator")
	return collection.Remove(bson.M{"_id": id})
}
