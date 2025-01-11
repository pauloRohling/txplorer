package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
)

type OperationRepository interface {
	Create(ctx context.Context, entity *operation.Operation) (*operation.Operation, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status operation.Status) (*operation.Operation, error)
}
