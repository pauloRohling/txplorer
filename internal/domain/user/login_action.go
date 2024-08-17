package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/password"
	"github.com/pauloRohling/txplorer/internal/domain/repository"
	"github.com/pauloRohling/txplorer/internal/domain/token"
	"github.com/pauloRohling/txplorer/internal/model"
	"time"
)

type LoginInput struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	TokenExpiration int64  `json:"-"`
}

type LoginOutput struct {
	AccessToken string `json:"accessToken"`
}

type LoginAction struct {
	userRepository     repository.UserRepository
	passwordComparator password.Comparator
	tokenGenerator     token.Generator
}

func NewLoginAction(userRepository repository.UserRepository, passwordComparator password.Comparator, tokenGenerator token.Generator) *LoginAction {
	return &LoginAction{
		userRepository:     userRepository,
		passwordComparator: passwordComparator,
		tokenGenerator:     tokenGenerator,
	}
}

func (action *LoginAction) Execute(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	user, err := action.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, model.NotFoundError("User not found")
	}

	if isEquals := action.passwordComparator.Compare(user.Password, input.Password); !isEquals {
		return nil, model.UnauthorizedError("Invalid credentials")
	}

	claims := action.generateClaims(user.ID, input.TokenExpiration)

	var accessToken string
	accessToken, err = action.tokenGenerator.Generate(claims)
	if err != nil {
		return nil, err
	}

	return action.fromToken(accessToken), nil
}

func (action *LoginAction) fromToken(accessToken string) *LoginOutput {
	return &LoginOutput{AccessToken: accessToken}
}

func (action *LoginAction) generateClaims(userId uuid.UUID, tokenExpiration int64) map[string]any {
	return map[string]any{
		"sub": userId.String(),
		"exp": time.Now().UTC().Add(time.Second * time.Duration(tokenExpiration)).Unix(),
	}
}
