package hasher

import "golang.org/x/crypto/bcrypt"

// HashPassword is a decorator for bcrypt.GenerateFromPassword.
// It returns a hashed string encrypted with the bcrypt algorithm.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// CheckPasswordHash is a decorator for bcrypt.CompareHashAndPassword.
// It returns true if the password match the hash or false if it's not.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
