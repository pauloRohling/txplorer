package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type OperationMapper interface {
	ToModel(operation store.Operation) *operation.Operation
}
