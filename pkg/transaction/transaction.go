package transaction

import (
	"context"
	"database/sql"
)

const contextKey = "x-transaction"

// Inject injects a transaction into the context
func Inject(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, contextKey, tx)
}

// Clean returns a context without the transaction
func Clean(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey, nil)
}

// FromContext returns the transaction from the context or nil
// if there is no transaction in the context
func FromContext(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(contextKey).(*sql.Tx); ok {
		return tx
	}
	return nil
}
