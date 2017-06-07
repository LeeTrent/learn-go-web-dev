package dao

import (
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/042_mongodb/10_hands-on/model"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

const (
	dbUrl = "mongodb://localhost"
	dbName = "go-web-dev-db"
	dbCollectionName = "users"
)

type UserDAO struct {
	dao *mgo.Session
}

func NewUserDAO() *UserDAO {

	mdbSession, err := mgo.Dial(dbUrl)

	if err != nil {
		panic(err)
	}

	return &UserDAO{dao: mdbSession}
}

func (ud UserDAO) Create(user model.User) (model.User) {

	user.Id = bson.NewObjectId()
	ud.dao.DB(dbName).C(dbCollectionName).Insert(user)
	return user
}

func (ud UserDAO) Retrieve(id string) (model.User, bool) {

	if !bson.IsObjectIdHex(id) {
		return model.User{}, false
	}

	oid := bson.ObjectIdHex(id)
	user := model.User{}

	if err := ud.dao.DB(dbName).C(dbCollectionName).FindId(oid).One(&user); err != nil {
		return user, false
	}
	return user, true
}

func (ud UserDAO) Delete(id string) (bool) {

	if !bson.IsObjectIdHex(id) {
		return false
	}

	oid := bson.ObjectIdHex(id)

	if err := ud.dao.DB(dbName).C(dbCollectionName).RemoveId(oid); err != nil {
		return false
	}
	return true
}