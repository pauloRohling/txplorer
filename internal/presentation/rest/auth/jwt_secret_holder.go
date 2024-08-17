package presentation

import (
	"github.com/go-chi/jwtauth/v5"
)

type SecretHolder interface {
	Get() *jwtauth.JWTAuth
}

type JwtSecretHolder struct {
	jwtAuth *jwtauth.JWTAuth
}

func NewJwtSecretHolder(secret string) *JwtSecretHolder {
	return &JwtSecretHolder{
		jwtAuth: jwtauth.New("HS256", []byte(secret), nil),
	}
}

func (holder *JwtSecretHolder) Get() *jwtauth.JWTAuth {
	return holder.jwtAuth
}
