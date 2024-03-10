package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword return hashed password from the given password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIfError(err)
	return string(hash)
}

// ComparePassword compare the given password with the hashed password
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
