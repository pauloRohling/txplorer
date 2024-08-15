package password

type Comparator interface {
	Compare(encodedPassword, password string) bool
}
