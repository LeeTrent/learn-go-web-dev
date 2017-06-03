package dao

import (
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/042_mongodb/06_hands-on/model"
	"github.com/satori/go.uuid"
)

type UserDAO struct {
	dao map[string]model.User
}

func NewUserDAO() *UserDAO {
	return &UserDAO { dao: make(map[string]model.User) }
}

func (ud UserDAO) Create(user model.User) (model.User) {
	user.Id = uuid.NewV4().String()
	ud.dao[user.Id] = user
	return ud.dao[user.Id]
}

func (ud UserDAO) Retrieve(id string) (model.User, bool) {
	user, found := ud.dao[id]
	return user, found
}

func (ud UserDAO) Delete(id string) (bool) {
	_, found := ud.dao[id]

	if found {
		delete (ud.dao, id)
	}
	return found
}