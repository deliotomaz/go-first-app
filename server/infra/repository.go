package infra

import (
	"../domain"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	pool           *mgosession.Pool
	Db             string
	CollectionName string
	Session        *mgo.Session
	collection     *mgo.Collection
}

type UserRepository MongoRepository
type WidgetRepository MongoRepository

//NewUserRepository create new repository for users
func NewUserRepository(p *mgosession.Pool, session *mgo.Session, db string) *UserRepository {
	return &UserRepository{
		pool:           p,
		CollectionName: "users",
		Session:        session,
		Db:             db,
		collection:     session.DB(db).C("users"),
	}
}

func (userRepo *UserRepository) Count() (int, error) {

	total, err := userRepo.collection.Count()
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (userRepo *UserRepository) FindById(id int64) (domain.User, error) {

	user := domain.User{}
	err := userRepo.collection.Find(bson.M{"id": id}).One(&user)
	return user, err
}
func (userRepo *UserRepository) List() (domain.Users, error) {

	users := domain.Users{}
	err := userRepo.collection.Find(bson.M{}).Sort("id").All(&users)
	return users, err
}

//NewUserRepository create new repository for users
func NewWidgetRepository(p *mgosession.Pool, session *mgo.Session, db string) *WidgetRepository {
	return &WidgetRepository{
		pool:           p,
		CollectionName: "widgets",
		Session:        session,
		Db:             db,
		collection:     session.DB(db).C("widgets"),
	}
}
func (widgetRepo *WidgetRepository) Count() (int, error) {

	total, err := widgetRepo.collection.Count()
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (widgetRepo *WidgetRepository) FindById(id int64) (domain.Widget, error) {

	widget := domain.Widget{}
	err := widgetRepo.collection.Find(bson.M{"id": id}).One(&widget)
	return widget, err
}
func (widgetRepo *WidgetRepository) List() (domain.Widgets, error) {

	widgets := domain.Widgets{}
	err := widgetRepo.collection.Find(bson.M{}).Sort("id").All(&widgets)
	return widgets, err
}
func (widgetRepo *WidgetRepository) AddUpdate(widget domain.Widget) error {

	//novo
	if widget.ID == 0 {
		total, err := widgetRepo.collection.Count()
		if err != nil {
			return err
		}
		widget.ID = int64(total) + 1
		err = widgetRepo.collection.Insert(&widget)
		return err
	} else {

		err := widgetRepo.collection.Update(bson.M{"id": widget.ID}, &widget)
		return err
	}

}
