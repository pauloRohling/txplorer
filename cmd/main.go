package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/mapper"
	"github.com/pauloRohling/txplorer/internal/persistance"
	"github.com/pauloRohling/txplorer/internal/presentation/rest"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/pkg/banner"
	"github.com/pauloRohling/txplorer/pkg/env"
	"github.com/pauloRohling/txplorer/pkg/envconfig"
	"github.com/pauloRohling/txplorer/pkg/graceful"
	tx "github.com/pauloRohling/txplorer/pkg/transaction"
	"log/slog"
	"os"
)

var environment env.Environment

func main() {
	banner.Show()
	_, _ = envconfig.Init(&environment)

	if err := environment.Validate(); err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}

	db, err := getDatabaseConnection()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}

	txManager := tx.NewPostgresTxManager(db)

	accountMapper := mapper.NewAccountMapper()
	transactionMapper := mapper.NewOperationMapper()

	accountRepository := persistance.NewAccountRepository(db, accountMapper)
	transactionRepository := persistance.NewTransactionRepository(db, transactionMapper)

	transferAction := operation.NewTransferAction(
		txManager,
		accountRepository,
		transactionRepository,
	)

	transactionService := operation.NewService(transferAction)

	transactionRouter := rest.NewOperationRouter(transactionService)

	httpServer := webserver.NewWebServer(environment.Server.Port, nil)
	gracefulShutdownCtx := graceful.Shutdown(&graceful.Params{
		OnStart:   func() { slog.Info("Graceful shutdown started. Waiting for active requests to complete") },
		OnTimeout: func() { slog.Error("Graceful shutdown timed out. Forcing exit.") },
		OnShutdown: func(timeoutCtx context.Context) {
			if err = httpServer.Shutdown(timeoutCtx); err != nil {
				slog.Error("Could not shutdown web server", "port", environment.Server.Port)
				os.Exit(-1)
			}
		},
	})

	httpServer.AddRoute(transactionRouter)

	slog.Info("Web server started listening on", "port", environment.Server.Port)
	if err = httpServer.Start(); err != nil {
		slog.Error("Could not start web server", "port", environment.Server.Port)
		os.Exit(-1)
	}

	<-gracefulShutdownCtx.Done()
	slog.Info("Graceful shutdown complete")
}

func getDatabaseConnectionString() string {
	ssl := "disable"
	if environment.Database.SSL {
		ssl = "require"
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		environment.Database.User,
		environment.Database.Password,
		environment.Database.Host,
		environment.Database.Port,
		environment.Database.Name,
		ssl,
	)
}

func getDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", getDatabaseConnectionString())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(environment.Database.Pool.MaxOpenConns)
	db.SetMaxIdleConns(environment.Database.Pool.MaxIdleConns)
	db.SetConnMaxLifetime(environment.Database.Pool.ConnMaxLifetime)
	db.SetConnMaxIdleTime(environment.Database.Pool.ConnMaxIdleTime)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
