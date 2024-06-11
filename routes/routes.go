package routes

import (
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/controllers"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/controllers/analysis"
	"github.com/gin-gonic/gin"
)

type Model interface {
	RegisterRoutes(*gin.Engine)
}

func SetupRouter(ginRouter *gin.Engine) {
	// List of all the possible routes services
	routeHandlers := []Model{
		&controllers.HealthCheckService{},
		&analysis.AnalyzeService{},

		// TODO: Register all the routing services here...
	}

	// Register each routes from the list of handlers
	for _, r := range routeHandlers {
		r.RegisterRoutes(ginRouter)
	}
}
