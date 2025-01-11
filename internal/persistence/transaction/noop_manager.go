package transaction

import (
	"context"
	"database/sql"
)

type NoopTxManager struct{}

func NewNoopTxManager() *NoopTxManager {
	return &NoopTxManager{}
}

func (manager *NoopTxManager) RunTransaction(ctx context.Context, fn func(context.Context) error) error {
	return fn(ctx)
}

func (manager *NoopTxManager) RunTransactionWithOptions(ctx context.Context, fn func(context.Context) error, _ *sql.TxOptions) error {
	return fn(ctx)
}
