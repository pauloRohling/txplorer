package presentation

type JwtGenerator struct {
	secretHolder SecretHolder
}

func NewJwtGenerator(secretHolder SecretHolder) *JwtGenerator {
	return &JwtGenerator{secretHolder: secretHolder}
}

func (generator *JwtGenerator) Generate(claims map[string]any) (string, error) {
	_, token, err := generator.secretHolder.Get().Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
