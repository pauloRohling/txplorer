package operation

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/domain/throw"
	"github.com/pauloRohling/txplorer/pkg/transaction"
	"time"
)

type DepositInput struct {
	AccountID   uuid.UUID `json:"accountId"`
	RequesterID uuid.UUID `json:"requesterId"`
	Amount      int64     `json:"amount"`
}

type DepositOutput struct {
	*operation.Operation
}

type DepositAction struct {
	accountRepository   account.Repository
	operationRepository operation.Repository
	transactionManager  transaction.Manager
}

func NewDepositAction(accountRepository account.Repository, operationRepository operation.Repository, transactionManager transaction.Manager) *DepositAction {
	return &DepositAction{
		accountRepository:   accountRepository,
		operationRepository: operationRepository,
		transactionManager:  transactionManager,
	}
}

func (action *DepositAction) Execute(ctx context.Context, input DepositInput) (*DepositOutput, error) {
	if input.Amount <= 0 {
		return nil, throw.ValidationError("Amount must be greater than 0 to make a deposit")
	}

	operationId, err := uuid.NewV7()
	if err != nil {
		return nil, throw.InternalError("Failed to generate operation id", err)
	}

	depositOperation := &operation.Operation{
		ID:            operationId,
		FromAccountID: input.AccountID,
		ToAccountID:   input.AccountID,
		Amount:        input.Amount,
		Type:          operation.DepositType.String(),
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     input.RequesterID,
		Status:        operation.PendingStatus,
	}

	newOperation, err := action.operationRepository.Create(ctx, depositOperation)
	if err != nil {
		return nil, err
	}

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		newOperation, err = action.updateBalance(ctx, input, operationId)
		return err
	})

	if err != nil {
		_, errOperation := action.operationRepository.UpdateStatus(ctx, operationId, operation.FailedStatus)
		if errOperation != nil {
			return nil, throw.InternalError("Failed to update operation status to FAILED", errOperation)
		}
		return nil, err
	}

	return &DepositOutput{Operation: newOperation}, nil
}

func (action *DepositAction) updateBalance(ctx context.Context, input DepositInput, operationId uuid.UUID) (*operation.Operation, error) {
	updatedAccount, err := action.accountRepository.AddBalanceById(ctx, input.AccountID, input.Amount)
	if err != nil {
		return nil, throw.InternalError("Failed to update account balance", err)
	}

	if updatedAccount.Balance < 0 {
		return nil, throw.ValidationError("Account balance is negative")
	}

	updatedStatus, err := action.operationRepository.UpdateStatus(ctx, operationId, operation.SuccessStatus)
	if err != nil {
		return nil, throw.InternalError("Failed to update operation status to SUCCESS", err)
	}

	return updatedStatus, nil
}
