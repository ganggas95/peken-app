package helper

import "golang.org/x/crypto/bcrypt"

type PasswordGenerator interface {
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
}

type PasswordGeneratorImpl struct {
}

// NewPasswordGenerator returns new PasswordGenerator.
func NewPasswordGenerator() *PasswordGeneratorImpl {
	return &PasswordGeneratorImpl{}
}

func (passGen *PasswordGeneratorImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (passGen *PasswordGeneratorImpl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
