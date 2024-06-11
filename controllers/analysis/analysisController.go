package analysis

import (
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/constants"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/errors"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/model/payload"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/model/response"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/utils"
	"github.com/gin-gonic/gin"
	"math"
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

	// Validate if there are more than 4 contracts
	if len(payload) > 4 {
		context.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrorResponse{
			ErrorCode: errors.ErrorCodeBadRequest,
			Message:   "too many option contracts",
		})
	}

	// TODO: Validate input fields.

	context.JSON(http.StatusOK, model_response.AnalysisResponse{
		XYValues:        calculateXYValues(payload),
		MaxProfit:       calculateMaxProfit(payload),
		MaxLoss:         calculateMaxLoss(payload),
		BreakEvenPoints: calculateBreakEvenPoints(payload),
	})
}

func calculateXYValues(contracts []model_payload.OptionsContract) []model_response.GraphPoint {
	priceRanges := getPriceRanges(contracts)
	profitLoss := getProfitLoss(contracts)

	graphPoints := make([]model_response.GraphPoint, len(priceRanges))

	for i := 0; i < len(priceRanges); i++ {
		graphPoints[i] = model_response.GraphPoint{
			X: utils.FormatFloatingNumberToCurrency(priceRanges[i]),
			Y: utils.FormatFloatingNumberToCurrency(profitLoss[i]),
		}
	}

	return graphPoints
}

func calculateMaxProfit(contracts []model_payload.OptionsContract) float64 {
	profitLoss := getProfitLoss(contracts)
	maxProfit := utils.FindMaximumFromFloatingNumbers(profitLoss)
	return utils.FormatFloatingNumberToCurrency(maxProfit)
}

func calculateMaxLoss(contracts []model_payload.OptionsContract) float64 {
	profitLoss := getProfitLoss(contracts)
	maxLoss := utils.FindMinimumFromFloatingNumbers(profitLoss)
	return utils.FormatFloatingNumberToCurrency(maxLoss)
}

func calculateBreakEvenPoints(contracts []model_payload.OptionsContract) []float64 {
	pricesRanges := getPriceRanges(contracts)
	profitLoss := getProfitLoss(contracts)
	breakEvenPoints := make([]float64, 0)
	for i := 1; i < len(pricesRanges); i++ {
		if (profitLoss[i-1] < 0 && profitLoss[i] > 0) || (profitLoss[i-1] > 0 && profitLoss[i] < 0) {
			breakEvenPoints = append(breakEvenPoints, utils.FormatFloatingNumberToCurrency(pricesRanges[i]))
		}
	}
	return breakEvenPoints
}

func getStrikePrices(contracts []model_payload.OptionsContract) []float64 {
	strikePrices := make([]float64, len(contracts))
	for index, contract := range contracts {
		strikePrices[index] = contract.StrikePrice
	}
	return strikePrices
}

func getPriceRanges(contracts []model_payload.OptionsContract) []float64 {
	strikePrices := getStrikePrices(contracts)
	minimumPrice := utils.FindMinimumFromFloatingNumbers(strikePrices) - constants.PriceRangeDifferentialFactor
	maximumPrice := utils.FindMaximumFromFloatingNumbers(strikePrices) + constants.PriceRangeDifferentialFactor
	priceRanges := utils.CreateLinearlySpacedFloatingNumbersArray(minimumPrice, maximumPrice, constants.NumberOfLinearlySpacedPriceRanges)
	return priceRanges
}

func getPremium(bid, ask float64) float64 {
	return utils.FormatFloatingNumberToCurrency((bid + ask) / 2)
}

func getProfitLoss(contracts []model_payload.OptionsContract) []float64 {
	prices := getPriceRanges(contracts)
	profitLoss := make([]float64, len(prices))

	for _, contract := range contracts {
		premium := getPremium(contract.Bid, contract.Ask)

		for i, price := range prices {
			switch contract.Type {
			case constants.ContractsOptionTypeCall:
				if contract.LongShort == constants.ContractsOptionLong {
					profitLoss[i] += math.Max(price-contract.StrikePrice, 0) - premium
				} else if contract.LongShort == constants.ContractsOptionShort {
					profitLoss[i] += premium - math.Max(price-contract.StrikePrice, 0)
				}

			case constants.ContractsOptionTypePut:
				if contract.LongShort == constants.ContractsOptionLong {
					profitLoss[i] += math.Max(contract.StrikePrice-price, 0) - premium
				} else if contract.LongShort == constants.ContractsOptionShort {
					profitLoss[i] += premium - math.Max(contract.StrikePrice-price, 0)
				}

			default:
			}
		}
	}

	return profitLoss
}
