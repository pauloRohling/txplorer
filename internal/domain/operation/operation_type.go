package operation

type Type string

const (
	TransferType Type = "TRANSFER"
	DepositType  Type = "DEPOSIT"
	WithdrawType Type = "WITHDRAW"
)

func (status Type) String() string {
	return string(status)
}
