package dao

import (
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/042_mongodb/08_hands-on/model"
	"github.com/satori/go.uuid"
	"os"
	"fmt"
	"encoding/json"
)

type UserDAO struct {
	dao map[string]model.User
}

func NewUserDAO() *UserDAO {

	//return &UserDAO { dao: make(map[string]model.User) }
	m := make(map[string]model.User)

	f, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return &UserDAO{dao: m}
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}
	return &UserDAO{dao: m}
}

func (ud UserDAO) Create(user model.User) (model.User) {
	user.Id = uuid.NewV4().String()
	ud.dao[user.Id] = user
	ud.persist()
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
		ud.persist()
	}
	return found
}

// Private method
func (ud UserDAO) persist() {

	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(ud.dao)
}