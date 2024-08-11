package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/mapper"
	"github.com/pauloRohling/txplorer/internal/persistance"
	"github.com/pauloRohling/txplorer/internal/presentation/rest"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/pkg/graceful"
	tx "github.com/pauloRohling/txplorer/pkg/transaction"
	"log/slog"
	"os"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/txplorer?sslmode=disable")
	if err != nil {
		panic(err)
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

	httpServer := webserver.NewWebServer(8080, nil)
	gracefulShutdownCtx := graceful.Shutdown(&graceful.Params{
		OnStart:   func() { slog.Info("Graceful shutdown started. Waiting for active requests to complete") },
		OnTimeout: func() { slog.Error("Graceful shutdown timed out. Forcing exit.") },
		OnShutdown: func(timeoutCtx context.Context) {
			slog.Info("Web server shutdown")
			if err = httpServer.Shutdown(timeoutCtx); err != nil {
				slog.Error("Could not shutdown web server", "port", "8080")
				os.Exit(-1)
			}
		},
	})

	httpServer.AddRoute(transactionRouter)

	slog.Info("Web server started listening on", "port", "8080")
	if err = httpServer.Start(); err != nil {
		slog.Error("Could not start web server", "port", "8080")
		os.Exit(-1)
	}

	<-gracefulShutdownCtx.Done()
	slog.Info("Graceful shutdown complete")
}
