package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/pauloRohling/txplorer/docs"
	"github.com/pauloRohling/txplorer/internal/application/account"
	"github.com/pauloRohling/txplorer/internal/application/operation"
	"github.com/pauloRohling/txplorer/internal/application/user"
	"github.com/pauloRohling/txplorer/internal/environment"
	accountPersitence "github.com/pauloRohling/txplorer/internal/persistence/account"
	operationPersistence "github.com/pauloRohling/txplorer/internal/persistence/operation"
	userPersistence "github.com/pauloRohling/txplorer/internal/persistence/user"
	accountPresentation "github.com/pauloRohling/txplorer/internal/presentation/rest/account"
	presentation "github.com/pauloRohling/txplorer/internal/presentation/rest/auth"
	operationPresentation "github.com/pauloRohling/txplorer/internal/presentation/rest/operation"
	userPresentation "github.com/pauloRohling/txplorer/internal/presentation/rest/user"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/pkg/crypto"
	"github.com/pauloRohling/txplorer/pkg/graceful"
	tx "github.com/pauloRohling/txplorer/pkg/transaction"
	"log/slog"
	"os"
	"time"
)

var start = time.Now()

//	@title			TxPlorer API
//	@version		1.0
//	@description	This is a transactional application that allows users to transfer funds between their accounts.
//	@contact.name	API Support
//	@contact.url	https://github.com/pauloRohling/txplorer
//	@license.name	MIT
//	@license.url	https://github.com/pauloRohling/txplorer/blob/master/LICENSE
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				"Authorization: Bearer <token>"
//	@tokenUrl					/users/login
//
// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	env := environment.Env()
	profile := environment.Profile()

	db, err := getDatabaseConnection(env)
	if err != nil {
		slog.Error("Could not get database connection", "error", err.Error())
		os.Exit(-1)
	}
	defer func(db *sql.DB) {
		if err = db.Close(); err != nil {
			slog.Error("Could not close database connection", "error", err.Error())
		}
	}(db)

	secretHolder := presentation.NewJwtSecretHolder(env.Security.Secret)
	passwordComparator := crypto.NewBcryptComparator()
	passwordEncoder := crypto.NewBcryptEncoder()
	tokenGenerator := presentation.NewJwtGenerator(secretHolder)
	txManager := tx.NewPostgresTxManager(db)

	accountMapper := accountPersitence.NewStoreMapper()
	operationMapper := operationPersistence.NewStoreMapper()
	userMapper := userPersistence.NewStoreMapper()

	accountRepository := accountPersitence.NewRepository(db, accountMapper)
	operationRepository := operationPersistence.NewRepository(db, operationMapper)
	userRepository := userPersistence.NewRepository(db, userMapper)

	createAccountAction := account.NewCreateAccountAction(accountRepository, userRepository, txManager, passwordEncoder)
	getAccountAction := account.NewGetAccountAction(accountRepository)
	depositAction := operation.NewDepositAction(accountRepository, operationRepository, txManager)
	loginAction := user.NewLoginAction(userRepository, passwordComparator, tokenGenerator, env.Security.TokenExpiration)
	transferAction := operation.NewTransferAction(accountRepository, operationRepository, txManager)
	withdrawAction := operation.NewWithdrawAction(accountRepository, operationRepository, txManager)

	accountService := account.NewFacadeService(createAccountAction, getAccountAction)
	operationService := operation.NewFacadeService(depositAction, transferAction, withdrawAction)
	userService := user.NewFacadeService(loginAction)

	accountRouter := accountPresentation.NewRestController(accountService, secretHolder)
	operationRouter := operationPresentation.NewRestController(operationService, secretHolder)
	userRouter := userPresentation.NewRestController(userService)

	httpServer := webserver.NewWebServer(env.Server.Port, nil)
	gracefulShutdownCtx := graceful.Shutdown(&graceful.Params{
		OnStart: func() {
			slog.Info("Graceful shutdown started. Waiting for active requests to complete")
		},
		OnTimeout: func() {
			slog.Error("Graceful shutdown timed out. Forcing exit.")
		},
		OnShutdown: func(timeoutCtx context.Context) {
			if err = httpServer.Shutdown(timeoutCtx); err != nil {
				slog.Error("Could not shutdown web server", "port", env.Server.Port)
				os.Exit(-1)
			}
		},
	})

	httpServer.AddRoute(accountRouter)
	httpServer.AddRoute(operationRouter)
	httpServer.AddRoute(userRouter)

	if profile == "dev" {
		httpServer.AddSwaggerRoute()
	}

	slog.Info("Web server started listening on", "port", env.Server.Port, "startup time", time.Since(start))
	if err = httpServer.Start(); err != nil {
		slog.Error("Could not start web server", "port", env.Server.Port)
		os.Exit(-1)
	}

	<-gracefulShutdownCtx.Done()
	slog.Info("Graceful shutdown complete")
}

func getDatabaseConnectionString(env environment.Config) string {
	ssl := "disable"
	if *env.Database.SSL {
		ssl = "require"
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		env.Database.User,
		env.Database.Password,
		env.Database.Host,
		env.Database.Port,
		env.Database.Name,
		ssl,
	)
}

func getDatabaseConnection(env environment.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", getDatabaseConnectionString(env))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(env.Database.Pool.MaxOpenConns)
	db.SetMaxIdleConns(env.Database.Pool.MaxIdleConns)
	db.SetConnMaxLifetime(env.Database.Pool.ConnMaxLifetime)
	db.SetConnMaxIdleTime(env.Database.Pool.ConnMaxIdleTime)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
