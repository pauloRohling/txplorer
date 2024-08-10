package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/model"
)

type OperationRepository interface {
	Create(ctx context.Context, entity *model.Operation) (*model.Operation, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status model.OperationStatus) (*model.Operation, error)
}
