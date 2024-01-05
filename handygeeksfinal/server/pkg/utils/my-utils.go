package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashMyPass(pass string) (string, error) {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost); err == nil {
		return string(hashedPassword), nil
	}
	return "", errors.New("Error hashing password")
}

func CompareMyPass(hashedPass string, plainPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass)); err == nil {
		return true
	}
	return false
}
