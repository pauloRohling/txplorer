package mapper

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistance/store"
)

type UserMapper interface {
	ToModel(user store.User) *model.User
}
