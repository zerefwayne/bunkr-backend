package utils

import "golang.org/x/crypto/bcrypt"

// CompareHashAndPassword takes in password of the user and the hashed password and compares them.
// Returns nil if passwords match.
func CompareHashAndPassword(password string, hash string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}

// GenerateHashFromPassword takes in a password string and hashes it
func GenerateHashFromPassword(password string) (string, error) {

	bytePassword := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)

	return string(hash), err

}
