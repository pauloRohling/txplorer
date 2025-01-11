package account

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
	"github.com/pauloRohling/txplorer/internal/persistence/transaction"
)

type Repository struct {
	db            *sql.DB
	accountMapper Mapper
}

func NewRepository(db *sql.DB, accountMapper Mapper) *Repository {
	return &Repository{
		db:            db,
		accountMapper: accountMapper,
	}
}

func (repository *Repository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(tx)
	}
	return store.New(repository.db)
}

func (repository *Repository) Create(ctx context.Context, userId uuid.UUID) (*account.Account, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	newAccount, err := repository.query(ctx).InsertAccount(ctx, store.InsertAccountParams{
		ID:     id,
		UserID: userId,
	})

	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(newAccount), nil
}

func (repository *Repository) AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*account.Account, error) {
	updatedAccount, err := repository.query(ctx).AddBalanceById(ctx, store.AddBalanceByIdParams{
		ID:      id,
		Balance: balance,
	})

	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(updatedAccount), nil
}

func (repository *Repository) GetById(ctx context.Context, id uuid.UUID) (*account.Account, error) {
	savedAccount, err := repository.query(ctx).GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(savedAccount), nil
}

func (repository *Repository) GetByUserId(ctx context.Context, userId uuid.UUID) (*account.Account, error) {
	savedAccount, err := repository.query(ctx).GetAccountByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(savedAccount), nil
}

var _ account.Repository = (*Repository)(nil)
