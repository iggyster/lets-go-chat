package tokenGenerator

import (
	"crypto/rand"
	"fmt"
)

func Generate(len int) (string, error) {
	token := make([]byte, len)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", token), nil
}
