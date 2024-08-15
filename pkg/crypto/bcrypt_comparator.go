package crypto

import "golang.org/x/crypto/bcrypt"

type BcryptComparator struct{}

func NewBcryptComparator() *BcryptComparator {
	return &BcryptComparator{}
}

func (b BcryptComparator) Compare(encodedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
	return err == nil
}
