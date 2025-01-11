package operation

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pauloRohling/throw"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/persistence/transaction"
	"time"
)

type TransferInput struct {
	FromAccountID uuid.UUID `json:"fromAccountId"`
	ToAccountID   uuid.UUID `json:"toAccountId"`
	RequesterID   uuid.UUID `json:"requesterId"`
	Amount        int64     `json:"amount"`
}

type TransferOutput struct {
	*operation.Operation
}

type TransferAction struct {
	accountRepository   account.Repository
	operationRepository operation.Repository
	transactionManager  transaction.Manager
}

func NewTransferAction(accountRepository account.Repository, operationRepository operation.Repository, transactionManager transaction.Manager) *TransferAction {
	return &TransferAction{
		accountRepository:   accountRepository,
		operationRepository: operationRepository,
		transactionManager:  transactionManager,
	}
}

func (action *TransferAction) Execute(ctx context.Context, input TransferInput) (*TransferOutput, error) {
	if input.FromAccountID == input.ToAccountID {
		return nil, throw.Validation().Msg("Cannot transfer to the same account")
	}

	if input.Amount <= 0 {
		return nil, throw.Validation().Msg("Invalid amount")
	}

	newAccount, err := action.accountRepository.GetById(ctx, input.FromAccountID)
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

	transferOperation := &operation.Operation{
		ID:            operationId,
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
		Type:          operation.TransferType.String(),
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     input.RequesterID,
		Status:        operation.PendingStatus,
	}

	newOperation, err := action.operationRepository.Create(ctx, transferOperation)
	if err != nil {
		return nil, err
	}

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		newOperation, err = action.updateBalances(ctx, input, operationId)
		return err
	})

	if err != nil {
		_, errOperation := action.operationRepository.UpdateStatus(ctx, operationId, operation.FailedStatus)
		if errOperation != nil {
			return nil, throw.Internal().Err(errOperation).Msg("Failed to update operation status to FAILED")
		}
		return nil, err
	}

	return &TransferOutput{Operation: newOperation}, nil
}

// UpdateBalances updates the balance of the sender and receiver accounts.
// Must be called inside a transaction.
func (action *TransferAction) updateBalances(ctx context.Context, input TransferInput, operationId uuid.UUID) (*operation.Operation, error) {
	fromAccount, err := action.accountRepository.AddBalanceById(ctx, input.FromAccountID, input.Amount*-1)
	if err != nil {
		return nil, throw.Internal().Err(err).Msgf("Failed to update account %s balance", input.FromAccountID)
	}

	if fromAccount.Balance < 0 {
		return nil, throw.Validation().Msgf("Account %s balance is negative", input.FromAccountID)
	}

	toAccount, err := action.accountRepository.AddBalanceById(ctx, input.ToAccountID, input.Amount)
	if err != nil {
		return nil, throw.Internal().Err(err).Msgf("Failed to update account %s balance", input.ToAccountID)
	}

	if toAccount.Balance < 0 {
		return nil, throw.Validation().Msg(fmt.Sprintf("Account %s balance is negative", input.ToAccountID))
	}

	updatedOperation, err := action.operationRepository.UpdateStatus(ctx, operationId, operation.SuccessStatus)
	if err != nil {
		return nil, throw.Internal().Err(err).Msg("Failed to update operation status to SUCCESS")
	}

	return updatedOperation, nil
}
