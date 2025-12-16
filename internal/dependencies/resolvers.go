package dependencies

import (
	"github.com/Sahil2k07/kakfa/internal/graphql/resolvers"
	"github.com/Sahil2k07/kakfa/internal/repositories"
	"github.com/Sahil2k07/kakfa/internal/services"
)

func Resolvers() *resolvers.Resolver {
	// Repositories
	authRepository := repositories.AuthRepository()

	// Services
	cryptoService := services.CryptoService()
	authService := services.AuthService(authRepository, cryptoService)

	return &resolvers.Resolver{
		AuthService: authService,
	}
}
