package account

type Status string

const (
	ActiveStatus   Status = "ACTIVE"
	InactiveStatus Status = "INACTIVE"
)

func (status Status) String() string {
	return string(status)
}
