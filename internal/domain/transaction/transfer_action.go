package transaction

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"xplorer/internal/model"
	tx "xplorer/pkg/transaction"
)

type TransferInput struct {
	FromAccountID uuid.UUID `json:"fromAccountId"`
	ToAccountID   uuid.UUID `json:"toAccountId"`
	Amount        int64     `json:"amount"`
}

type TransferOutput struct {
	*model.Transaction
}

type TransferAction struct {
	transactionManager    tx.Manager
	accountRepository     AccountRepository
	transactionRepository TransactionRepository
}

func NewTransferAction(transactionManager tx.Manager, accountRepository AccountRepository, transactionRepository TransactionRepository) *TransferAction {
	return &TransferAction{
		transactionManager:    transactionManager,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (action *TransferAction) Execute(ctx context.Context, input TransferInput) (*TransferOutput, error) {
	if input.FromAccountID == input.ToAccountID {
		return nil, fmt.Errorf("[TransferAction] cannot transfer to the same account: %s", input.ToAccountID)
	}

	if input.Amount <= 0 {
		return nil, fmt.Errorf("[TransferAction] invalid amount: %d", input.Amount)
	}

	newTransaction := model.NewTransaction(input.FromAccountID, input.ToAccountID, input.Amount)
	transaction, err := action.transactionRepository.Create(ctx, newTransaction)
	if err != nil {
		return nil, fmt.Errorf("[TransferAction] failed to create transaction: %w", err)
	}

	transactionId := transaction.ID

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		transaction, err = action.updateBalances(ctx, input, transactionId)
		return err
	})

	if err != nil {
		_, errTransaction := action.transactionRepository.UpdateStatus(ctx, transactionId, model.TransactionStatusFailed)
		if errTransaction != nil {
			return nil, fmt.Errorf("[TransferAction] failed to update transaction %s status to FAILED: %w", transactionId, errTransaction)
		}
		return nil, err
	}

	return &TransferOutput{Transaction: transaction}, nil
}

// UpdateBalances updates the balance of the sender and receiver accounts.
// Must be called inside a transaction.
func (action *TransferAction) updateBalances(ctx context.Context, input TransferInput, transactionId uuid.UUID) (*model.Transaction, error) {
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

	transaction, err := action.transactionRepository.UpdateStatus(ctx, transactionId, model.TransactionStatusSuccess)
	if err != nil {
		return nil, fmt.Errorf("[TransferAction] failed to update transaction %s status to SUCCESS: %w", transactionId, err)
	}

	return transaction, nil
}
