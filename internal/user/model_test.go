package user

import (
	"github.com/google/uuid"
	"github.com/iggyster/lets-go-chat/pkg/hasher"
	"github.com/iggyster/lets-go-chat/pkg/tokengenerator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_RevokeToken(t *testing.T) {
	usr := New("test", "pass")
	token := tokengenerator.Generate(16)

	usr.SetToken(token)
	usr.RevokeToken()

	assert.Empty(t, usr.Token)
}

func TestUser_SetToken(t *testing.T) {
	usr := New("test", "pass")
	token := tokengenerator.Generate(16)

	usr.SetToken(token)

	assert.Equal(t, token, usr.Token)
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

			assert.NoError(t, err)
			assert.Equal(t, input.username, usr.Username)
			assert.True(t, hasher.CheckPasswordHash(input.password, usr.Password))
			assert.Empty(t, usr.Token)
		})
	}
}
