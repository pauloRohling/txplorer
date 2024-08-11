package envconfig

type Profile string

const (
	Dev  Profile = "dev"
	Prod Profile = "prod"
)

func NewProfile(profile string) Profile {
	switch profile {
	case Prod.String():
		return Prod
	default:
		return Dev
	}
}

func (p Profile) String() string {
	return string(p)
}
