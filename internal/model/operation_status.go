package model

type OperationStatus string

const (
	OperationStatusPending OperationStatus = "PENDING"
	OperationStatusSuccess OperationStatus = "SUCCESS"
	OperationStatusFailed  OperationStatus = "FAILED"
)

func (status OperationStatus) String() string {
	return string(status)
}
