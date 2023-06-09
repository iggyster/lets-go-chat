package user

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/iggyster/lets-go-chat/pkg/hasher"
	"github.com/iggyster/lets-go-chat/pkg/tokengen"
)

func TestUser_Token(t *testing.T) {
	usr := New("test", "pass")
	token := tokengen.New(16)

	usr.SetToken(token)
	if 0 == len(usr.Token) {
		t.Errorf("faield to set a token")
	}

	usr.RevokeToken()

	if 0 != len(usr.Token) {
		t.Errorf("fail to revoke a token")
	}
}

func TestUser_Activation(t *testing.T) {
	usr := New("test", "pass")

	if usr.IsActivated() {
		t.Errorf("user must be inactive by default")
	}

	usr.Activate()
	if !usr.IsActivated() {
		t.Errorf("fail to activate user")
	}

	usr.Deactivate()
	if usr.IsActivated() {
		t.Errorf("fail to deactivate user")
	}
}

func TestNew(t *testing.T) {
	data := []struct {
		name, username, password string
	}{
		{name: "Empty", username: "", password: ""},
		{name: "Single", username: "u", password: "p"},
		{name: "ABC-only", username: "username", password: "password"},
		{name: "Numbers-only", username: "123123", password: "123123"},
		{name: "Mixed", username: "test@123", password: "test@123"},
	}

	for _, input := range data {
		t.Run(input.name, func(t *testing.T) {
			usr := New(input.username, input.password)
			_, err := uuid.Parse(usr.Id)

			if err != nil {
				t.Errorf("fail to generate uuid for a user")
			}

			if input.username != usr.Username {
				t.Errorf("fail create a user with user name %v", input.username)
			}

			if !hasher.CheckPasswordHash(input.password, usr.Password) {
				t.Errorf("user password has changed during the user intitiation")
			}

			if 0 != len(usr.Token) {
				t.Error("user must have nnot toekn on initialization")
			}
		})
	}
}

func ExampleNew() {
	usr := New("test", "secret")

	fmt.Println(usr.Username)

	// Output: test
}
