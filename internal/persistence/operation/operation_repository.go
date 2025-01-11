package operation

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
	"github.com/pauloRohling/txplorer/internal/persistence/transaction"
)

type Repository struct {
	db              *sql.DB
	operationMapper Mapper
}

func NewRepository(db *sql.DB, operationMapper Mapper) *Repository {
	return &Repository{
		db:              db,
		operationMapper: operationMapper,
	}
}

func (repository *Repository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(tx)
	}
	return store.New(repository.db)
}

func (repository *Repository) Create(ctx context.Context, entity *operation.Operation) (*operation.Operation, error) {
	transactionEntity, err := repository.query(ctx).InsertOperation(ctx, store.InsertOperationParams{
		ID:            entity.ID,
		FromAccountID: entity.FromAccountID,
		ToAccountID:   entity.ToAccountID,
		Amount:        entity.Amount,
		Type:          entity.Type,
		CreatedBy:     entity.CreatedBy,
	})

	if err != nil {
		return nil, err
	}

	return repository.operationMapper.ToModel(transactionEntity), nil
}

func (repository *Repository) UpdateStatus(ctx context.Context, id uuid.UUID, status operation.Status) (*operation.Operation, error) {
	transactionEntity, err := repository.query(ctx).UpdateOperationStatus(ctx, store.UpdateOperationStatusParams{
		ID:     id,
		Status: status.String(),
	})

	if err != nil {
		return nil, err
	}

	return repository.operationMapper.ToModel(transactionEntity), nil
}

var _ operation.Repository = (*Repository)(nil)
