package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type OperationMapper struct {
}

func NewOperationMapper() *OperationMapper {
	return &OperationMapper{}
}

func (mapper *OperationMapper) ToModel(savedOperation store.Operation) *operation.Operation {
	return &operation.Operation{
		ID:            savedOperation.ID,
		FromAccountID: savedOperation.FromAccountID,
		ToAccountID:   savedOperation.ToAccountID,
		Amount:        savedOperation.Amount,
		Type:          savedOperation.Type,
		CreatedAt:     savedOperation.CreatedAt,
		CreatedBy:     savedOperation.CreatedBy,
		Status:        operation.Status(savedOperation.Status),
	}
}
