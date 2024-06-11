package analysis

import (
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/errors"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/model/payload"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/model/response"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func analyzeOptionsContracts(context *gin.Context) {
	payload := make([]model_payload.OptionsContract, 0)
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrorResponse{
			ErrorCode: errors.ErrorCodeBadRequest,
			Message:   utils.ParseValidationError(err),
		})
	}

	context.JSON(http.StatusOK, payload)
}

func calculateXYValues(contracts []model_payload.OptionsContract) []model_response.GraphPoint {
	// Your code here
	return nil
}

func calculateMaxProfit(contracts []model_payload.OptionsContract) float64 {
	// Your code here
	return 0
}

func calculateMaxLoss(contracts []model_payload.OptionsContract) float64 {
	// Your code here
	return 0
}

func calculateBreakEvenPoints(contracts []model_payload.OptionsContract) []float64 {
	// Your code here
	return nil
}
