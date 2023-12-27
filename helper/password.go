package helper

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtils interface {
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
}

type PasswordUtilsImpl struct {
}

// NewPasswordUtils returns new PasswordUtils.
func NewPasswordUtils() *PasswordUtilsImpl {
	return &PasswordUtilsImpl{}
}

func (passGen *PasswordUtilsImpl) HashPassword(password string) (string, error) {
	// HashPassword generates a hashed password from the given string.
	//
	// It takes a string parameter `password` which is the password to be hashed.
	// It returns a string which is the hashed password and an error if any.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (passGen *PasswordUtilsImpl) CheckPasswordHash(password string, hash string) bool {
	// CheckPasswordHash checks if a password matches its corresponding hash.
	//
	// password: the password to be checked.
	// hash: the hash to be compared against the password.
	// returns: a boolean indicating whether the password matches the hash.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
