package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type OperationMapper struct {
}

func NewOperationMapper() *OperationMapper {
	return &OperationMapper{}
}

func (mapper *OperationMapper) ToModel(operation store.Operation) *domain.Operation {
	return &domain.Operation{
		ID:            operation.ID,
		FromAccountID: operation.FromAccountID,
		ToAccountID:   operation.ToAccountID,
		Amount:        operation.Amount,
		Type:          operation.Type,
		CreatedAt:     operation.CreatedAt,
		CreatedBy:     operation.CreatedBy,
		Status:        domain.OperationStatus(operation.Status),
	}
}
