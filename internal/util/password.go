package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword - Password hashing and comparison utilities
func HashPassword(password string) (string, error) {
	hshPassw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hshPassw), nil
}

// CheckPassword compares a plaintext password with a hashed password
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
