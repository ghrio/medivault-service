package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Hashes Password
func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

// Compares Password
func comparePassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

// encrypts data
func Encrypt[T any](data []byte) ([]byte, error) {
	return data, nil
}

// decrypts data
func Decrypt[T any](data []byte) ([]byte, error) {
	return data, nil
}
