package operation

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/repository"
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/pkg/transaction"
	"time"
)

type WithdrawInput struct {
	AccountID   uuid.UUID `json:"accountId"`
	RequesterID uuid.UUID `json:"requesterId"`
	Amount      int64     `json:"amount"`
}

type WithdrawOutput struct {
	*model.Operation
}

type WithdrawAction struct {
	accountRepository   repository.AccountRepository
	operationRepository repository.OperationRepository
	transactionManager  transaction.Manager
}

func NewWithdrawAction(accountRepository repository.AccountRepository, operationRepository repository.OperationRepository, transactionManager transaction.Manager) *WithdrawAction {
	return &WithdrawAction{
		accountRepository:   accountRepository,
		operationRepository: operationRepository,
		transactionManager:  transactionManager,
	}
}

func (action *WithdrawAction) Execute(ctx context.Context, input WithdrawInput) (*WithdrawOutput, error) {
	if input.Amount <= 0 {
		return nil, model.ValidationError("Amount must be greater than 0 to make a withdrawal")
	}

	operationId, err := uuid.NewV7()
	if err != nil {
		return nil, model.InternalError("Failed to generate operation id", err)
	}

	withdrawOperation := &model.Operation{
		ID:            operationId,
		FromAccountID: input.AccountID,
		ToAccountID:   input.AccountID,
		Amount:        input.Amount,
		Type:          model.OperationTypeWithdraw.String(),
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     input.RequesterID,
		Status:        model.OperationStatusPending,
	}

	operation, err := action.operationRepository.Create(ctx, withdrawOperation)
	if err != nil {
		return nil, err
	}

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		operation, err = action.updateBalance(ctx, input, operationId)
		return err
	})

	if err != nil {
		_, errOperation := action.operationRepository.UpdateStatus(ctx, operationId, model.OperationStatusFailed)
		if errOperation != nil {
			return nil, model.InternalError("Failed to update operation status to FAILED", errOperation)
		}
		return nil, err
	}

	return &WithdrawOutput{Operation: operation}, nil
}

func (action *WithdrawAction) updateBalance(ctx context.Context, input WithdrawInput, operationId uuid.UUID) (*model.Operation, error) {
	account, err := action.accountRepository.AddBalanceById(ctx, input.AccountID, input.Amount*-1)
	if err != nil {
		return nil, model.InternalError("Failed to update account balance", err)
	}

	if account.Balance < 0 {
		return nil, model.ValidationError("Account balance is negative")
	}

	operation, err := action.operationRepository.UpdateStatus(ctx, operationId, model.OperationStatusSuccess)
	if err != nil {
		return nil, model.InternalError("Failed to update operation status to SUCCESS", err)
	}

	return operation, nil
}
