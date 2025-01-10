package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (mapper *UserMapper) ToModel(user store.User) *domain.User {
	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
