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

type OperationRepository struct {
	db              *sql.DB
	operationMapper mapper.OperationMapper
}

func NewTransactionRepository(db *sql.DB, operationMapper mapper.OperationMapper) *OperationRepository {
	return &OperationRepository{
		db:              db,
		operationMapper: operationMapper,
	}
}

func (repository *OperationRepository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(tx)
	}
	return store.New(repository.db)
}

func (repository *OperationRepository) Create(ctx context.Context, entity *model.Operation) (*model.Operation, error) {
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

func (repository *OperationRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status model.OperationStatus) (*model.Operation, error) {
	transactionEntity, err := repository.query(ctx).UpdateOperationStatus(ctx, store.UpdateOperationStatusParams{
		ID:     id,
		Status: status.String(),
	})

	if err != nil {
		return nil, err
	}

	return repository.operationMapper.ToModel(transactionEntity), nil
}
