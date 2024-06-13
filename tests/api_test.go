package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/api"
	model_payload "github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/model/payload"
	"github.com/gin-gonic/gin"
	"math"
	"os"
	"time"
)

const (
	FilePathValidOptionsContracts           = "../testdata/testdata.json"
	FilePathMissingRequiredField            = "../testdata/testdata_with_missing_required_field.json"
	FilePathTooManyOptionsContracts         = "../testdata/testdata_with_too_many_options_contracts.json"
	FilePathInvalidOptionsContractType      = "../testdata/testdata_with_invalid_type.json"
	FilePathInvalidOptionsContractLongShort = "../testdata/testdata_with_invalid_long_short.json"

	ValidOptionsContractsExpirationDate = "2025-12-17T00:00:00Z"
)

type TestImplementation struct {
	api.API
	consumed bool
}

func (apiTest *TestImplementation) NewTestServer() *gin.Engine {
	if apiTest.consumed {
		panic("API implementation was already used")
	}
	apiTest.consumed = true
	return apiTest.API.CreateServer()
}

func initializeTestServer() *gin.Engine {
	repositoryProvider := mockDependencyProvider{}
	return (&TestImplementation{
		API: api.NewBuilder().
			WithRepositoryProvider(repositoryProvider).
			Finalize(),
	}).NewTestServer()
}

func getMockedTestData(filePath string) []model_payload.OptionsContract {
	optionsContracts := make([]model_payload.OptionsContract, 4)
	raw, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(raw, &optionsContracts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return optionsContracts
}

func getAverageNoOfDays(contracts []model_payload.OptionsContract) int {
	totalNoOfDays := 0
	for _, contract := range contracts {
		totalNoOfDays += int(math.Round(contract.ExpirationDate.Sub(time.Now()).Hours() / 24))
	}
	return int(math.Round(float64(totalNoOfDays) / float64(len(contracts))))
}
