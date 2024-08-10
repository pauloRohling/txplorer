package mapper

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistance/store"
)

type OperationMapper struct {
}

func NewOperationMapper() *OperationMapper {
	return &OperationMapper{}
}

func (mapper *OperationMapper) ToModel(operation store.Operation) *model.Operation {
	return &model.Operation{
		ID:            operation.ID,
		FromAccountID: operation.FromAccountID,
		ToAccountID:   operation.ToAccountID,
		Amount:        operation.Amount,
		Type:          operation.Type,
		CreatedAt:     operation.CreatedAt,
		CreatedBy:     operation.CreatedBy,
		Status:        operation.Status,
	}
}
