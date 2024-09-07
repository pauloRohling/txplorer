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

type AccountRepository struct {
	db            *sql.DB
	accountMapper mapper.AccountMapper
}

func NewAccountRepository(db *sql.DB, accountMapper mapper.AccountMapper) *AccountRepository {
	return &AccountRepository{
		db:            db,
		accountMapper: accountMapper,
	}
}

func (repository *AccountRepository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(tx)
	}
	return store.New(repository.db)
}

func (repository *AccountRepository) Create(ctx context.Context, userId uuid.UUID) (*model.Account, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	account, err := repository.query(ctx).InsertAccount(ctx, store.InsertAccountParams{
		ID:     id,
		UserID: userId,
	})

	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(account), nil
}

func (repository *AccountRepository) AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*model.Account, error) {
	account, err := repository.query(ctx).AddBalanceById(ctx, store.AddBalanceByIdParams{
		ID:      id,
		Balance: balance,
	})

	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(account), nil
}

func (repository *AccountRepository) GetById(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	account, err := repository.query(ctx).GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(account), nil
}
