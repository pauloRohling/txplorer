package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type UserMapper interface {
	ToModel(user store.User) *domain.User
}
