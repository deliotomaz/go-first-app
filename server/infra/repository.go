package infra

import (
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	pool       *mgosession.Pool
	Db         string
	Collection string
	Session    *mgo.Session
}

type UserRepository MongoRepository
type WidgetRepository MongoRepository

//NewUserRepository create new repository for users
func NewUserRepository(p *mgosession.Pool, session *mgo.Session, db string) *UserRepository {
	return &UserRepository{
		pool:       p,
		Collection: "Users",
		Session:    session,
		Db:         db,
	}
}

func (userRepo *UserRepository) Count() (int, error) {

	coll := userRepo.Session.DB(userRepo.Db).C(userRepo.Collection)
	total, e := coll.Count()
	if e != nil {
		return 0, e
	}
	return total, nil
}
