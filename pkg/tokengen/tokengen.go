package tokengen

import (
	"crypto/rand"
	"fmt"
)

// New generates a new string token according to the length.
func New(len int) string {
	t := make([]byte, len/2)
	rand.Read(t)

	return fmt.Sprintf("%x", t)
}
