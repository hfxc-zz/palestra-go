package collaborator

import (
	"palestra-go/pkg/entity"

	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	pool *mgosession.Pool
	db   string
}

func NewMongoRepository(p *mgosession.Pool, db string) *MongoRepository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}

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

func (r *MongoRepository) Search(query string) ([]*entity.Collaborator, error) {
	return []*entity.Collaborator{}, nil
}

func (r *MongoRepository) Save(c *entity.Collaborator) (bson.ObjectId, error) {
	session := r.pool.Session(nil)
	collection := session.DB(r.db).C("collaborator")
	err := collection.Insert(c)
	if err != nil {
		return bson.ObjectId(0), err
	}
	return c.ID, nil
}

func (r *MongoRepository) Delete(id bson.ObjectId) error {
	return nil
}
