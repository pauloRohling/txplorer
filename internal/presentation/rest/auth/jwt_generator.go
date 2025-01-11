package presentation

import "github.com/pauloRohling/txplorer/internal/domain/token"

type JwtGenerator struct {
	secretHolder SecretHolder
}

func NewJwtGenerator(secretHolder SecretHolder) *JwtGenerator {
	return &JwtGenerator{secretHolder: secretHolder}
}

func (generator *JwtGenerator) Generate(claims map[string]any) (string, error) {
	_, newToken, err := generator.secretHolder.Get().Encode(claims)
	if err != nil {
		return "", err
	}
	return newToken, nil
}

var _ token.Generator = (*JwtGenerator)(nil)
