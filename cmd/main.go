package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
	"time"
	"xplorer/internal/domain/transaction"
	"xplorer/internal/mapper"
	"xplorer/internal/persistance"
	"xplorer/internal/presentation/json"
	"xplorer/internal/presentation/rest"
	tx "xplorer/pkg/transaction"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/xplorer?sslmode=disable")
	if err != nil {
		panic(err)
	}

	txManager := tx.NewPostgresTxManager(db)

	accountMapper := mapper.NewAccountMapper()
	transactionMapper := mapper.NewTransactionMapper()

	accountRepository := persistance.NewAccountRepository(db, accountMapper)
	transactionRepository := persistance.NewTransactionRepository(db, transactionMapper)

	transferAction := transaction.NewTransferAction(txManager, accountRepository, transactionRepository)

	transactionService := transaction.NewService(transferAction)

	transactionRouter := rest.NewTransactionRouter(transactionService)

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.Timeout(30 * time.Second))

	router.Route("/api/v1/transactions", func(r chi.Router) {
		r.Post("/", json.Endpoint(transactionRouter.Transfer, http.StatusCreated))
	})

	slog.Info("Web server started listening on", "port", "8080")
	if err = http.ListenAndServe(":8080", router); err != nil {
		slog.Error("Could not initialize web server", "port", "8080")
		os.Exit(-1)
	}

}
