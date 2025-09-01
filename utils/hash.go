package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

func ComparePasswords(retrievedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(retrievedPassword), []byte(password))
	return err == nil
}
