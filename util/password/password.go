package password

import (
	"golang.org/x/crypto/bcrypt"
)

// VerifyPassword compares password and the hashed password
func VerifyPassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

func VerifySecret(passwordHash, password string) error {
	if password == passwordHash {
		return nil
	} else {
		return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	}
}

// HashPassword creates a bcrypt password hash
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 3)
}
