package analysis

import "github.com/gin-gonic/gin"

type AnalyzeService struct{}

func (analyzeService *AnalyzeService) RegisterRoutes(router *gin.Engine) {
	router.POST("/analyze", analyzeOptionsContracts)
}
