package hashpass

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword)
}

func VerifyPassword(hashedPassword string, enteredPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
}
