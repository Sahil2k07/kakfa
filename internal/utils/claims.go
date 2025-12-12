package utils

import (
	"context"
	"time"

	errz "github.com/Sahil2k07/kakfa/internal/errors"
)

type UserClaims struct {
	ID        uint
	Email     string
	UserName  string
	ExpiresAt *time.Time
}

type userCtxKey struct{}

var UserCtxKey = userCtxKey{}

func GetUserClaims(ctx context.Context) (*UserClaims, error) {
	raw := ctx.Value(UserCtxKey)
	if raw == nil {
		return nil, errz.NewUnauthorized("no user claims in context")
	}

	claims, ok := raw.(*UserClaims)
	if !ok {
		return nil, errz.NewUnauthorized("invalid claims type in context")
	}

	return claims, nil
}
