package tokengenerator

import (
	"crypto/rand"
	"fmt"
)

func Generate(len int) string {
	token := make([]byte, len)
	rand.Read(token)

	return fmt.Sprintf("%x", token)
}
