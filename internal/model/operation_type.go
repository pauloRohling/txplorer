package model

type OperationType string

const (
	OperationTypeTransfer OperationType = "TRANSFER"
	OperationTypeDeposit  OperationType = "DEPOSIT"
	OperationTypeWithdraw OperationType = "WITHDRAW"
)

func (status OperationType) String() string {
	return string(status)
}
