package persistance

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistance/mapper"
	"github.com/pauloRohling/txplorer/internal/persistance/store"
	"github.com/pauloRohling/txplorer/pkg/transaction"
)

type UserRepository struct {
	db         *sql.DB
	userMapper mapper.UserMapper
}

func NewUserRepository(db *sql.DB, userMapper mapper.UserMapper) *UserRepository {
	return &UserRepository{
		db:         db,
		userMapper: userMapper,
	}
}

func (repository *UserRepository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(tx)
	}
	return store.New(repository.db)
}

func (repository *UserRepository) Create(ctx context.Context, name string, email string, password string) (*model.User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	user, err := repository.query(ctx).InsertUser(ctx, store.InsertUserParams{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return repository.userMapper.ToModel(user), nil
}
