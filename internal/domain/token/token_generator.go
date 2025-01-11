package token

type Generator interface {
	Generate(claims map[string]any) (string, error)
}
