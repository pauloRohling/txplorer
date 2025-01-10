package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type OperationMapper interface {
	ToModel(operation store.Operation) *domain.Operation
}
