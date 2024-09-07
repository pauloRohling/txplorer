package middleware

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"net/http"
)

const (
	UserIdContextKey = "txplorer-auth-userId"
)

type Claims map[string]any

func Authenticator(jwtAuth *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, claims, err := jwtauth.FromContext(r.Context())

			if err != nil {
				json.WriteError(w, model.UnauthorizedError("Could not find token in context", err))
				return
			}

			if token == nil || jwt.Validate(token, jwtAuth.ValidateOptions()...) != nil {
				json.WriteError(w, model.UnauthorizedError("Token is not valid"))
				return
			}

			ctx, err := createAuthContext(r.Context(), claims)
			if err != nil {
				json.WriteError(w, err)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func createAuthContext(ctx context.Context, claims Claims) (context.Context, error) {
	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, model.UnauthorizedError("Could not find 'sub' claim")
	}

	userId, err := uuid.Parse(sub)
	if err != nil {
		return nil, model.UnauthorizedError("Could not parse 'sub' claim as uuid")
	}

	ctx = context.WithValue(ctx, UserIdContextKey, userId)
	return ctx, nil
}

func GetUserId(ctx context.Context) (uuid.UUID, error) {
	userId, ok := ctx.Value(UserIdContextKey).(uuid.UUID)
	if !ok {
		return uuid.UUID{}, model.UnauthorizedError("could not get 'userId' from context")
	}
	return userId, nil
}
