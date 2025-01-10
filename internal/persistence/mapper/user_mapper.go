package mapper

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type UserMapper interface {
	ToModel(user store.User) *model.User
}
