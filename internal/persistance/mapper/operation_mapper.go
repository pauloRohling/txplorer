package mapper

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistance/store"
)

type OperationMapper interface {
	ToModel(operation store.Operation) *model.Operation
}
