package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain"
)

type OperationRepository interface {
	Create(ctx context.Context, entity *domain.Operation) (*domain.Operation, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status domain.OperationStatus) (*domain.Operation, error)
}
