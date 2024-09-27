package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/pauloRohling/txplorer/docs"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/domain/user"
	"github.com/pauloRohling/txplorer/internal/mapper"
	"github.com/pauloRohling/txplorer/internal/persistance"
	presentation "github.com/pauloRohling/txplorer/internal/presentation/rest/auth"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/router"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/pkg/banner"
	"github.com/pauloRohling/txplorer/pkg/crypto"
	"github.com/pauloRohling/txplorer/pkg/env"
	"github.com/pauloRohling/txplorer/pkg/envconfig"
	"github.com/pauloRohling/txplorer/pkg/graceful"
	tx "github.com/pauloRohling/txplorer/pkg/transaction"
	"log/slog"
	"os"
	"time"
)

var (
	start       = time.Now()
	environment env.Environment
)

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
	banner.Show()
	profile, _ := envconfig.Init(&environment)

	if err := environment.Validate(); err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}

	db, err := getDatabaseConnection()
	if err != nil {
		slog.Error("Could not get database connection", "error", err.Error())
		os.Exit(-1)
	}
	defer func(db *sql.DB) {
		if err = db.Close(); err != nil {
			slog.Error("Could not close database connection", "error", err.Error())
		}
	}(db)

	secretHolder := presentation.NewJwtSecretHolder(environment.Security.Secret)
	passwordComparator := crypto.NewBcryptComparator()
	passwordEncoder := crypto.NewBcryptEncoder()
	tokenGenerator := presentation.NewJwtGenerator(secretHolder)
	txManager := tx.NewPostgresTxManager(db)

	accountMapper := mapper.NewAccountMapper()
	operationMapper := mapper.NewOperationMapper()
	userMapper := mapper.NewUserMapper()

	accountRepository := persistance.NewAccountRepository(db, accountMapper)
	operationRepository := persistance.NewOperationRepository(db, operationMapper)
	userRepository := persistance.NewUserRepository(db, userMapper)

	createAccountAction := account.NewCreateAccountAction(accountRepository, userRepository, txManager, passwordEncoder)
	getAccountAction := account.NewGetAccountAction(accountRepository)
	depositAction := operation.NewDepositAction(accountRepository, operationRepository, txManager)
	loginAction := user.NewLoginAction(userRepository, passwordComparator, tokenGenerator, environment.Security.TokenExpiration)
	transferAction := operation.NewTransferAction(accountRepository, operationRepository, txManager)
	withdrawAction := operation.NewWithdrawAction(accountRepository, operationRepository, txManager)

	accountService := account.NewService(createAccountAction, getAccountAction)
	operationService := operation.NewService(depositAction, transferAction, withdrawAction)
	userService := user.NewService(loginAction)

	accountRouter := router.NewAccountRouter(accountService, secretHolder)
	operationRouter := router.NewOperationRouter(operationService, secretHolder)
	userRouter := router.NewUserRouter(userService)

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

	httpServer.AddRoute(accountRouter)
	httpServer.AddRoute(operationRouter)
	httpServer.AddRoute(userRouter)

	if profile == envconfig.Dev {
		httpServer.AddSwaggerRoute()
	}

	slog.Info("Web server started listening on", "port", environment.Server.Port, "startup time", time.Since(start))
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
