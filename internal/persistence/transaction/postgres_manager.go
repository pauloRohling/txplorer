package transaction

import (
	"context"
	"database/sql"
)

type PostgresTxManager struct {
	db *sql.DB
}

func NewPostgresTxManager(db *sql.DB) *PostgresTxManager {
	return &PostgresTxManager{db: db}
}

func (manager *PostgresTxManager) RunTransaction(ctx context.Context, fn func(context.Context) error) error {
	return manager.RunTransactionWithOptions(ctx, fn, nil)
}

func (manager *PostgresTxManager) RunTransactionWithOptions(ctx context.Context, fn func(context.Context) error, options *sql.TxOptions) error {
	tx, err := manager.db.BeginTx(ctx, options)
	if err != nil {
		return err
	}

	contextWithTx := Inject(ctx, tx)
	if err = fn(contextWithTx); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}
		return err
	}

	return nil
}
