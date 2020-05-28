package utils

import "golang.org/x/crypto/bcrypt"

// CompareHashAndPassword ...
func CompareHashAndPassword(password string, hash string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}

// GenerateHashFromPassword ...
func GenerateHashFromPassword(password string) (string, error) {

	bytePassword := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)

	return string(hash), err

}
