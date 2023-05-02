package user

type Repo interface {
	FindByUsername(username string) *User
	IsExists(username string) bool
	Save(user *User)
}

type InMemoryRepo struct {
	store map[string]*User
}

var Repository Repo = &InMemoryRepo{store: map[string]*User{}}

func (repo *InMemoryRepo) FindByUsername(username string) *User {
	user, ok := repo.store[username]
	if !ok {
		return &User{}
	}

	return user
}

func (repo *InMemoryRepo) IsExists(username string) bool {
	_, ok := repo.store[username]

	return ok
}

func (repo *InMemoryRepo) Save(user *User) {
	repo.store[user.Username] = user
}
