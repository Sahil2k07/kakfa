package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	errz "github.com/Sahil2k07/kakfa/internal/errors"
	"github.com/Sahil2k07/kakfa/internal/utils"
)

func AuthDirectiveMiddleware() graphql.FieldMiddleware {
	return func(ctx context.Context, next graphql.Resolver) (any, error) {

		fieldCtx := graphql.GetFieldContext(ctx)
		if fieldCtx == nil {
			return next(ctx)
		}

		if ctx.Value(utils.UserCtxKey) != nil {
			return next(ctx)
		}

		for _, d := range fieldCtx.Field.Definition.Directives {
			if d.Name == "public" {
				return next(ctx) // allow public field
			}
		}

		return nil, errz.NewUnauthorized("unauthorized: missing token")
	}
}
