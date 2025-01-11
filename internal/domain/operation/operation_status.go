package operation

type Status string

const (
	PendingStatus Status = "PENDING"
	SuccessStatus Status = "SUCCESS"
	FailedStatus  Status = "FAILED"
)

func (status Status) String() string {
	return string(status)
}
