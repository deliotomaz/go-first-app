package infra

import (
	"github.com/juju/mgosession"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	pool *mgosession.Pool
	db   string
}

type UserRepository MongoRepository
type WidgetRepository MongoRepository

//NewUserRepository create new repository for users
func NewUserRepository(p *mgosession.Pool, db string) *UserRepository {
	return &UserRepository{
		pool: p,
		db:   db,
	}
}

func (userRepo *UserRepository) Count() (*int64, error) {
	return nil, nil
}
