package user

import (
	"fmt"
	"testing"
)

func TestNewRepo(t *testing.T) {
	var repo UserRepo = NewRepo()

	_, ok := repo.(UserRepo)
	if !ok {
		t.Errorf("fail to create new repo")
	}
}

func TestInMemoryRepo_FindByUser_InvalidType(t *testing.T) {
	var repo = NewRepo()

	key := "username"
	repo.Store(key, "not a user")

	_, err := repo.FindByUsername(key)
	if err == nil {
		t.Errorf("fail to find user type by username")
	}
}

func TestInMemoryRepo_FindByToken(t *testing.T) {
	var repo UserRepo = NewRepo()

	usr := New("test", "pass")
	usr.SetToken("token")

	repo.Save(usr)

	match := repo.FindByToken("token")
	if match == nil {
		t.Errorf("fail to find user by token")
	}
}

func TestInMemoryRepo_FindActivated(t *testing.T) {
	var repo UserRepo = NewRepo()

	usr := New("test", "pass")
	usr.Activate()

	repo.Save(usr)

	res := repo.FindActivated()
	if len(res) != 1 {
		t.Errorf("fail to fin activated users")
	}
}

func TestInMemoryRepo_IsExists(t *testing.T) {
	var repo UserRepo = NewRepo()

	usr := New("test", "pass")
	repo.Save(usr)

	if !repo.IsExists("test") {
		t.Errorf("fail to check if user is exist")
	}
}

func ExampleUserRepo_FindByUsername() {
	var repo UserRepo = NewRepo()

	name := "test"
	usr := New(name, "pass")
	repo.Save(usr)

	_, err := repo.FindByUsername("unknown")
	if err != nil {
		fmt.Println("User not found")
	}

	_, err = repo.FindByUsername(name)
	if err == nil {
		fmt.Println("User has been found")
	}

	// Output:
	// User not found
	// User has been found
}

func ExampleUserRepo_FindByToken() {
	var repo UserRepo = NewRepo()

	usr := New("test", "pass")
	usr.SetToken("token")
	repo.Save(usr)

	match := repo.FindByToken("notatoken")
	if match == nil {
		fmt.Println("User not found")
	}

	match = repo.FindByToken("token")
	if match != nil {
		fmt.Println("User has been found")
	}

	// Output:
	// User not found
	// User has been found
}

func BenchmarkFindByUsername(b *testing.B) {
	var repo UserRepo = NewRepo()

	usr := New("test", "pass")
	repo.Save(usr)

	b.ResetTimer()
	for num := 0; num < b.N; num++ {
		repo.FindByUsername("test")
	}
}

func BenchmarkFindByToken(b *testing.B) {
	var repo UserRepo = NewRepo()

	usr := New("test", "pass")
	usr.SetToken("token")
	repo.Save(usr)

	b.ResetTimer()
	for num := 0; num < b.N; num++ {
		repo.FindByToken("token")
	}
}

func BenchmarkIsExists(b *testing.B) {
	var repo UserRepo = NewRepo()

	usr := New("test", "pass")
	repo.Save(usr)

	b.ResetTimer()
	for num := 0; num < b.N; num++ {
		repo.IsExists("test")
	}
}
