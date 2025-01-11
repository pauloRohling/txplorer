package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain/user"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type UserMapper interface {
	ToModel(user store.User) *user.User
}
