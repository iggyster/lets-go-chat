package user

import (
	"fmt"
	"sync"
)

type Repo interface {
	FindByUsername(username string) (*User, error)
	IsExists(username string) bool
	Save(user *User)
}

type InMemoryRepo struct {
	store sync.Map
}

var Repository Repo = &InMemoryRepo{store: sync.Map{}}

func (repo *InMemoryRepo) FindByUsername(username string) (*User, error) {
	val, ok := repo.store.Load(username)
	if !ok {
		return nil, fmt.Errorf("no user found")
	}

	usr, ok := val.(*User)
	if !ok {
		return nil, fmt.Errorf("invalid user type %T", usr)
	}

	return usr, nil
}

func (repo *InMemoryRepo) IsExists(username string) bool {
	_, ok := repo.store.Load(username)

	return ok
}

func (repo *InMemoryRepo) Save(usr *User) {
	repo.store.Store(usr.Username, usr)
}
