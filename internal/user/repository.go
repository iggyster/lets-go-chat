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
	for _, user := range repo.store {
		if user.Username == username {
			return user
		}
	}

	return &User{}
}

func (repo *InMemoryRepo) IsExists(username string) bool {
	return repo.FindByUsername(username) != nil
}

func (repo *InMemoryRepo) Save(user *User) {
	repo.store[user.Id] = user
}
