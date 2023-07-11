package user

import "github.com/google/wire"

var InMemoryUserRepoProviderSet = wire.NewSet(
	ProvideInMemoryUserRepo,
	wire.Bind(new(UserRepo), new(*InMemoryUserRepo)),
)

func ProvideInMemoryUserRepo() *InMemoryUserRepo {
	return new(InMemoryUserRepo)
}
