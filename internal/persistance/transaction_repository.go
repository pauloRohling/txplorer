package persistance

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistance/mapper"
	"github.com/pauloRohling/txplorer/internal/persistance/store"
	"github.com/pauloRohling/txplorer/pkg/transaction"
)

type TransactionRepository struct {
	db                *sql.DB
	transactionMapper mapper.TransactionMapper
}

func NewTransactionRepository(db *sql.DB, transactionMapper mapper.TransactionMapper) *TransactionRepository {
	return &TransactionRepository{
		db:                db,
		transactionMapper: transactionMapper,
	}
}

func (repository *TransactionRepository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(tx)
	}
	return store.New(repository.db)
}

func (repository *TransactionRepository) Create(ctx context.Context, entity *model.Transaction) (*model.Transaction, error) {
	transactionEntity, err := repository.query(ctx).InsertTransaction(ctx, store.InsertTransactionParams{
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
	transactionEntity, err := repository.query(ctx).UpdateTransactionStatus(ctx, store.UpdateTransactionStatusParams{
		ID:     id,
		Status: status.String(),
	})

	if err != nil {
		return nil, err
	}

	return repository.transactionMapper.ToModel(transactionEntity), nil
}
