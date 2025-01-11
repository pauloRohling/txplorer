package operation

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/throw"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/persistence/transaction"
	"time"
)

type WithdrawInput struct {
	AccountID   uuid.UUID `json:"accountId"`
	RequesterID uuid.UUID `json:"requesterId"`
	Amount      int64     `json:"amount"`
}

type WithdrawOutput struct {
	*operation.Operation
}

type WithdrawAction struct {
	accountRepository   account.Repository
	operationRepository operation.Repository
	transactionManager  transaction.Manager
}

func NewWithdrawAction(accountRepository account.Repository, operationRepository operation.Repository, transactionManager transaction.Manager) *WithdrawAction {
	return &WithdrawAction{
		accountRepository:   accountRepository,
		operationRepository: operationRepository,
		transactionManager:  transactionManager,
	}
}

func (action *WithdrawAction) Execute(ctx context.Context, input WithdrawInput) (*WithdrawOutput, error) {
	if input.Amount <= 0 {
		return nil, throw.Validation().Msg("Amount must be greater than 0 to make a withdrawal")
	}

	newAccount, err := action.accountRepository.GetById(ctx, input.AccountID)
	if err != nil {
		return nil, throw.Internal().Err(err).Msg("Failed to get account")
	}

	if newAccount.UserID != input.RequesterID {
		return nil, throw.Unauthorized().Msg("You are not authorized to make this withdrawal")
	}

	operationId, err := uuid.NewV7()
	if err != nil {
		return nil, throw.Internal().Err(err).Msg("Failed to generate operation id")
	}

	withdrawOperation := &operation.Operation{
		ID:            operationId,
		FromAccountID: input.AccountID,
		ToAccountID:   input.AccountID,
		Amount:        input.Amount,
		Type:          operation.WithdrawType.String(),
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     input.RequesterID,
		Status:        operation.PendingStatus,
	}

	newOperation, err := action.operationRepository.Create(ctx, withdrawOperation)
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
			return nil, throw.Internal().Err(errOperation).Msg("Failed to update operation status to FAILED")
		}
		return nil, err
	}

	return &WithdrawOutput{Operation: newOperation}, nil
}

func (action *WithdrawAction) updateBalance(ctx context.Context, input WithdrawInput, operationId uuid.UUID) (*operation.Operation, error) {
	newAccount, err := action.accountRepository.AddBalanceById(ctx, input.AccountID, input.Amount*-1)
	if err != nil {
		return nil, throw.Internal().Err(err).Msg("Failed to update account balance")
	}

	if newAccount.Balance < 0 {
		return nil, throw.Validation().Msg("Account balance is negative")
	}

	updatedOperation, err := action.operationRepository.UpdateStatus(ctx, operationId, operation.SuccessStatus)
	if err != nil {
		return nil, throw.Internal().Err(err).Msg("Failed to update operation status to SUCCESS")
	}

	return updatedOperation, nil
}
