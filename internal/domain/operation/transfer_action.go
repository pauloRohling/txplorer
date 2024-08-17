package operation

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/repository"
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/pkg/transaction"
	"time"
)

type TransferInput struct {
	FromAccountID uuid.UUID `json:"fromAccountId"`
	ToAccountID   uuid.UUID `json:"toAccountId"`
	RequesterID   uuid.UUID `json:"requesterId"`
	Amount        int64     `json:"amount"`
}

type TransferOutput struct {
	*model.Operation
}

type TransferAction struct {
	accountRepository   repository.AccountRepository
	operationRepository repository.OperationRepository
	transactionManager  transaction.Manager
}

func NewTransferAction(accountRepository repository.AccountRepository, operationRepository repository.OperationRepository, transactionManager transaction.Manager) *TransferAction {
	return &TransferAction{
		accountRepository:   accountRepository,
		operationRepository: operationRepository,
		transactionManager:  transactionManager,
	}
}

func (action *TransferAction) Execute(ctx context.Context, input TransferInput) (*TransferOutput, error) {
	if input.FromAccountID == input.ToAccountID {
		return nil, model.ValidationError("Cannot transfer to the same account")
	}

	if input.Amount <= 0 {
		return nil, model.ValidationError("Invalid amount")
	}

	operationId, err := uuid.NewV7()
	if err != nil {
		return nil, model.InternalError("Failed to generate operation id", err)
	}

	transferOperation := &model.Operation{
		ID:            operationId,
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
		Type:          model.OperationTypeTransfer.String(),
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     input.RequesterID,
		Status:        model.OperationStatusPending,
	}

	operation, err := action.operationRepository.Create(ctx, transferOperation)
	if err != nil {
		return nil, err
	}

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		operation, err = action.updateBalances(ctx, input, operationId)
		return err
	})

	if err != nil {
		_, errOperation := action.operationRepository.UpdateStatus(ctx, operationId, model.OperationStatusFailed)
		if errOperation != nil {
			return nil, model.InternalError("Failed to update operation status to FAILED", errOperation)
		}
		return nil, err
	}

	return &TransferOutput{Operation: operation}, nil
}

// UpdateBalances updates the balance of the sender and receiver accounts.
// Must be called inside a transaction.
func (action *TransferAction) updateBalances(ctx context.Context, input TransferInput, operationId uuid.UUID) (*model.Operation, error) {
	fromAccount, err := action.accountRepository.AddBalanceById(ctx, input.FromAccountID, input.Amount*-1)
	if err != nil {
		return nil, model.InternalError(fmt.Sprintf("Failed to update account %s balance", input.FromAccountID), err)
	}

	if fromAccount.Balance < 0 {
		return nil, model.ValidationError(fmt.Sprintf("Account %s balance is negative", input.FromAccountID))
	}

	toAccount, err := action.accountRepository.AddBalanceById(ctx, input.ToAccountID, input.Amount)
	if err != nil {
		return nil, model.InternalError(fmt.Sprintf("Failed to update account %s balance", input.ToAccountID), err)
	}

	if toAccount.Balance < 0 {
		return nil, model.ValidationError(fmt.Sprintf("Account %s balance is negative", input.ToAccountID))
	}

	operation, err := action.operationRepository.UpdateStatus(ctx, operationId, model.OperationStatusSuccess)
	if err != nil {
		return nil, model.InternalError("Failed to update operation status to SUCCESS", err)
	}

	return operation, nil
}
