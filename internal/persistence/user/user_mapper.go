package user

import (
	"github.com/pauloRohling/txplorer/internal/domain/user"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type Mapper interface {
	ToModel(user store.User) *user.User
}

type StoreMapper struct {
}

func NewStoreMapper() *StoreMapper {
	return &StoreMapper{}
}

func (mapper *StoreMapper) ToModel(savedUser store.User) *user.User {
	return &user.User{
		ID:        savedUser.ID,
		Name:      savedUser.Name,
		Email:     savedUser.Email,
		Password:  savedUser.Password,
		CreatedAt: savedUser.CreatedAt,
		UpdatedAt: savedUser.UpdatedAt,
	}
}

var _ Mapper = (*StoreMapper)(nil)
