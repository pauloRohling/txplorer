package operation

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/repository"
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/pkg/transaction"
	"time"
)

type DepositInput struct {
	AccountID   uuid.UUID `json:"accountId"`
	RequesterID uuid.UUID `json:"requesterId"`
	Amount      int64     `json:"amount"`
}

type DepositOutput struct {
	*model.Operation
}

type DepositAction struct {
	accountRepository   repository.AccountRepository
	operationRepository repository.OperationRepository
	transactionManager  transaction.Manager
}

func NewDepositAction(accountRepository repository.AccountRepository, operationRepository repository.OperationRepository, transactionManager transaction.Manager) *DepositAction {
	return &DepositAction{
		accountRepository:   accountRepository,
		operationRepository: operationRepository,
		transactionManager:  transactionManager,
	}
}

func (action *DepositAction) Execute(ctx context.Context, input DepositInput) (*DepositOutput, error) {
	if input.Amount <= 0 {
		return nil, model.ValidationError("Amount must be greater than 0 to make a deposit")
	}

	operationId, err := uuid.NewV7()
	if err != nil {
		return nil, model.InternalError("Failed to generate operation id", err)
	}

	depositOperation := &model.Operation{
		ID:            operationId,
		FromAccountID: input.AccountID,
		ToAccountID:   input.AccountID,
		Amount:        input.Amount,
		Type:          model.OperationTypeDeposit.String(),
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     input.RequesterID,
		Status:        model.OperationStatusPending,
	}

	operation, err := action.operationRepository.Create(ctx, depositOperation)
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

	return &DepositOutput{Operation: operation}, nil
}

func (action *DepositAction) updateBalance(ctx context.Context, input DepositInput, operationId uuid.UUID) (*model.Operation, error) {
	account, err := action.accountRepository.AddBalanceById(ctx, input.AccountID, input.Amount)
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
