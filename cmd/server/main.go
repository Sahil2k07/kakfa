package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Sahil2k07/kakfa/internal/configs"
	"github.com/Sahil2k07/kakfa/internal/connections"
	"github.com/Sahil2k07/kakfa/internal/dependencies"
	"github.com/Sahil2k07/kakfa/internal/graphql/directives"
	"github.com/Sahil2k07/kakfa/internal/graphql/generated"
	"github.com/Sahil2k07/kakfa/internal/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs := configs.LoadConfig()

	connections.ConnectWDB()
	connections.ConnectKafkaWriter()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: configs.Server.Origins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Use(middlewares.JWTContext()) // JWT check if it's there or not

	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers:  dependencies.Resolvers(), // Dependency Injection for the application
		Directives: generated.DirectiveRoot{},
	}))

	gqlHandler.AroundFields(directives.AuthDirectiveMiddleware()) // If not Auth-Token, check for the @public directive

	e.POST("/graphql", echo.WrapHandler(gqlHandler))
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))

	e.Logger.Fatal(e.Start(configs.Server.Port))
}
