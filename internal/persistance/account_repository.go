package persistance

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"xplorer/internal/model"
	"xplorer/internal/persistance/queries"
	"xplorer/pkg/transaction"
)

type AccountRepository struct {
	db            *sql.DB
	accountMapper AccountMapper
}

func NewAccountRepository(db *sql.DB, accountMapper AccountMapper) *AccountRepository {
	return &AccountRepository{
		db:            db,
		accountMapper: accountMapper,
	}
}

func (repository *AccountRepository) query(ctx context.Context) *queries.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return queries.New(tx)
	}
	return queries.New(repository.db)
}

func (repository *AccountRepository) Create(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	account, err := repository.query(ctx).InsertAccount(ctx, id)

	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(account), nil
}

func (repository *AccountRepository) FindById(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	account, err := repository.query(ctx).GetAccountById(ctx, id)

	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(account), nil
}

func (repository *AccountRepository) AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*model.Account, error) {
	account, err := repository.query(ctx).AddBalanceById(ctx, queries.AddBalanceByIdParams{
		ID:      id,
		Balance: balance,
	})

	if err != nil {
		return nil, err
	}

	return repository.accountMapper.ToModel(account), nil
}
