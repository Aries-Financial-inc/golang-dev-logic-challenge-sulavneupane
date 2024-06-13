package api

import (
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/api/provider"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/constants"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/routes"
	"github.com/gin-gonic/gin"
)

type defaultRepositoryProvider struct {
}

var _ provider.RepositoryProvider = &defaultRepositoryProvider{}

func GetRepositories() provider.RepositoryProvider {
	return &defaultRepositoryProvider{}
}

type Builder interface {
	WithRepositoryProvider(provider.RepositoryProvider) Builder
	Finalize() API
}

type builder struct {
	provider provider.RepositoryProvider
}

func (b *builder) WithRepositoryProvider(provider provider.RepositoryProvider) Builder {
	b.provider = provider
	return b
}

func (b *builder) Finalize() API {
	if b.provider == nil {
		panic("No Dependency Provider Set")
	}
	return &Implementation{
		repositoryProvider: b.provider,
	}
}

func NewBuilder() Builder {
	return &builder{}
}

type API interface {
	ListenAndServe() error
	Repositories() provider.RepositoryProvider
	CreateServer() *gin.Engine
}

type Implementation struct {
	repositoryProvider provider.RepositoryProvider
}

func (api *Implementation) Repositories() provider.RepositoryProvider {
	return api.repositoryProvider
}

func (api *Implementation) providerMiddleware(c *gin.Context) {
	providers := map[string]interface{}{
		constants.DependencyProviderContextKey: api.Repositories(),
	}
	for key, value := range providers {
		c.Set(key, value)
	}
	c.Next()
}

func (api *Implementation) CreateServer() *gin.Engine {
	ginRouter := gin.Default()

	ginRouter.Use(api.providerMiddleware)

	// Register all routes
	routes.SetupRouter(ginRouter)

	return ginRouter
}

func (api *Implementation) ListenAndServe() error {
	ginRouter := api.CreateServer()
	return ginRouter.Run(":8080")
}
