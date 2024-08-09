package persistance

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"xplorer/internal/model"
	"xplorer/internal/persistance/queries"
	"xplorer/pkg/transaction"
)

type TransactionRepository struct {
	db                *sql.DB
	transactionMapper TransactionMapper
}

func NewTransactionRepository(db *sql.DB, transactionMapper TransactionMapper) *TransactionRepository {
	return &TransactionRepository{
		db:                db,
		transactionMapper: transactionMapper,
	}
}

func (repository *TransactionRepository) query(ctx context.Context) *queries.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return queries.New(tx)
	}
	return queries.New(repository.db)
}

func (repository *TransactionRepository) Create(ctx context.Context, entity *model.Transaction) (*model.Transaction, error) {
	transactionEntity, err := repository.query(ctx).InsertTransaction(ctx, queries.InsertTransactionParams{
		ID:            entity.ID,
		FromAccountID: entity.FromAccountID,
		ToAccountID:   entity.ToAccountID,
		Amount:        entity.Amount,
		Timestamp:     *entity.Timestamp,
		Status:        entity.Status,
	})

	if err != nil {
		return nil, err
	}

	return repository.transactionMapper.ToModel(transactionEntity), nil
}

func (repository *TransactionRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status model.TransactionStatus) (*model.Transaction, error) {
	transactionEntity, err := repository.query(ctx).UpdateTransactionStatus(ctx, queries.UpdateTransactionStatusParams{
		ID:     id,
		Status: status.String(),
	})

	if err != nil {
		return nil, err
	}

	return repository.transactionMapper.ToModel(transactionEntity), nil
}
