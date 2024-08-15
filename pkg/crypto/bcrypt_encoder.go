package crypto

import "golang.org/x/crypto/bcrypt"

type BcryptEncoder struct{}

func NewBcryptEncoder() *BcryptEncoder {
	return &BcryptEncoder{}
}

func (b BcryptEncoder) Encode(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
