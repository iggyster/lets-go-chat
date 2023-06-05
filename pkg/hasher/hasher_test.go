package hasher

import (
	"testing"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	pass, err := password.Generate(21, 7, 7, false, true)
	if err != nil {
		t.Errorf("failed to generate a new password")
	}

	actual, _ := HashPassword(pass)
	err = bcrypt.CompareHashAndPassword([]byte(actual), []byte(pass))

	if err != nil {
		t.Errorf("hash result doesn't match bcrypt generation with the default cost")
	}
}

func TestHashPasswordExceedLimit(t *testing.T) {
	pass, err := password.Generate(73, 7, 7, false, true)
	if err != nil {
		t.Errorf("Failed to generate a new password")
	}

	_, err = HashPassword(pass)

	if err == nil {
		t.Errorf("hash function can't apply passwords longer the 72 characters")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	pass, err := password.Generate(21, 7, 7, false, true)
	if err != nil {
		t.Errorf("Failed to generate a new password")
	}

	hash, _ := HashPassword(pass)
	actual := CheckPasswordHash(pass, hash)

	if !actual {
		t.Errorf("password check fails to match the password with its hash")
	}
}
