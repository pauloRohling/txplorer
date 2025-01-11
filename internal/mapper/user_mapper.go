package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain/user"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (mapper *UserMapper) ToModel(savedUser store.User) *user.User {
	return &user.User{
		ID:        savedUser.ID,
		Name:      savedUser.Name,
		Email:     savedUser.Email,
		Password:  savedUser.Password,
		CreatedAt: savedUser.CreatedAt,
		UpdatedAt: savedUser.UpdatedAt,
	}
}
