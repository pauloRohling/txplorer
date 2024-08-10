package operation

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/repository"
	"github.com/pauloRohling/txplorer/internal/model"
	tx "github.com/pauloRohling/txplorer/pkg/transaction"
	"time"
)

type TransferInput struct {
	FromAccountID uuid.UUID `json:"fromAccountId"`
	ToAccountID   uuid.UUID `json:"toAccountId"`
	Amount        int64     `json:"amount"`
	RequesterID   uuid.UUID `json:"requesterId"`
}

type TransferOutput struct {
	*model.Operation
}

type TransferAction struct {
	transactionManager  tx.Manager
	accountRepository   repository.AccountRepository
	operationRepository repository.OperationRepository
}

func NewTransferAction(
	transactionManager tx.Manager,
	accountRepository repository.AccountRepository,
	operationRepository repository.OperationRepository,
) *TransferAction {
	return &TransferAction{
		transactionManager:  transactionManager,
		accountRepository:   accountRepository,
		operationRepository: operationRepository,
	}
}

func (action *TransferAction) Execute(ctx context.Context, input TransferInput) (*TransferOutput, error) {
	if input.FromAccountID == input.ToAccountID {
		return nil, fmt.Errorf("[TransferAction] cannot transfer to the same account: %s", input.ToAccountID)
	}

	if input.Amount <= 0 {
		return nil, fmt.Errorf("[TransferAction] invalid amount: %d", input.Amount)
	}

	operationId, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("[TransferAction] failed to generate operation id: %w", err)
	}

	transferOperation := &model.Operation{
		ID:            operationId,
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
		Type:          model.OperationTypeTransfer.String(),
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     input.RequesterID,
		Status:        model.OperationStatusPending.String(),
	}

	operation, err := action.operationRepository.Create(ctx, transferOperation)
	if err != nil {
		return nil, fmt.Errorf("[TransferAction] failed to create operation: %w", err)
	}

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		operation, err = action.updateBalances(ctx, input, operationId)
		return err
	})

	if err != nil {
		_, errOperation := action.operationRepository.UpdateStatus(ctx, operationId, model.OperationStatusFailed)
		if errOperation != nil {
			return nil, fmt.Errorf("[TransferAction] failed to update operation %s status to FAILED: %w", operationId, errOperation)
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
		return nil, fmt.Errorf("[TransferAction] failed to update account %s balance: %w", input.FromAccountID, err)
	}

	if fromAccount.Balance < 0 {
		return nil, fmt.Errorf("[TransferAction] account %s balance is negative: %d", input.FromAccountID, fromAccount.Balance)
	}

	toAccount, err := action.accountRepository.AddBalanceById(ctx, input.ToAccountID, input.Amount)
	if err != nil {
		return nil, fmt.Errorf("[TransferAction] failed to update account %s balance: %w", input.ToAccountID, err)
	}

	if toAccount.Balance < 0 {
		return nil, fmt.Errorf("[TransferAction] account %s balance is negative: %d", input.ToAccountID, toAccount.Balance)
	}

	operation, err := action.operationRepository.UpdateStatus(ctx, operationId, model.OperationStatusSuccess)
	if err != nil {
		return nil, fmt.Errorf("[TransferAction] failed to update operation %s status to SUCCESS: %w", operationId, err)
	}

	return operation, nil
}
