package password

type Encoder interface {
	Encode(string) (string, error)
}
