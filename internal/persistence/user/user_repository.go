package user

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pauloRohling/throw"
	"github.com/pauloRohling/txplorer/internal/domain/user"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
	"github.com/pauloRohling/txplorer/internal/persistence/transaction"
)

type Repository struct {
	db         *sql.DB
	userMapper Mapper
}

func NewRepository(db *sql.DB, userMapper Mapper) *Repository {
	return &Repository{
		db:         db,
		userMapper: userMapper,
	}
}

func (repository *Repository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(tx)
	}
	return store.New(repository.db)
}

func (repository *Repository) Create(ctx context.Context, name string, email string, password string) (*user.User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, throw.Internal().Err(err).Msg("Failed to generate user id")
	}

	savedUser, err := repository.query(ctx).InsertUser(ctx, store.InsertUserParams{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, throw.Internal().Err(err).Msg("Failed to create user")
	}

	return repository.userMapper.ToModel(savedUser), nil
}

func (repository *Repository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	savedUser, err := repository.query(ctx).FindUserByEmail(ctx, email)
	if err != nil {
		return nil, throw.NotFound().Err(err).Msg("User not found")
	}

	return repository.userMapper.ToModel(savedUser), nil
}

var _ user.Repository = (*Repository)(nil)
