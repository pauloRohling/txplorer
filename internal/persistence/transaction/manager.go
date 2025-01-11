package transaction

import (
	"context"
	"database/sql"
)

type Manager interface {
	RunTransaction(ctx context.Context, fn func(context.Context) error) error
	RunTransactionWithOptions(ctx context.Context, fn func(context.Context) error, options *sql.TxOptions) error
}
