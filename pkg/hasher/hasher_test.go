package hasher

import (
	"github.com/sethvargo/go-password/password"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	pass, err := password.Generate(21, 7, 7, false, true)
	if err != nil {
		t.Errorf("Failed to generate a new password")
	}

	actual, _ := HashPassword(pass)
	err = bcrypt.CompareHashAndPassword([]byte(actual), []byte(pass))

	assert.Nil(t, err, "Hash result doesn't match bcrypt generation with the default cost")
}

func TestHashPasswordExceedLimit(t *testing.T) {
	pass, err := password.Generate(73, 7, 7, false, true)
	if err != nil {
		t.Errorf("Failed to generate a new password")
	}

	_, err = HashPassword(pass)

	assert.Error(t, err, "Hash function can't apply passwords longer the 72 characters")
}

func TestCheckPasswordHash(t *testing.T) {
	pass, err := password.Generate(21, 7, 7, false, true)
	if err != nil {
		t.Errorf("Failed to generate a new password")
	}

	hash, _ := HashPassword(pass)
	actual := CheckPasswordHash(pass, hash)

	assert.Truef(t, actual, "Password check fails to match the password with its hash")
}
