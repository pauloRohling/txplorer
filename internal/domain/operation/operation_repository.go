package operation

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, entity *Operation) (*Operation, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status Status) (*Operation, error)
}
