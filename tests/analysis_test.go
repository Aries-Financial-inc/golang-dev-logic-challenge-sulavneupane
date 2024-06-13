package tests

import (
	"encoding/json"
	"fmt"
	model_response "github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/model/response"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
	healthCheckHappyPathTest(t)
}

func TestOptionsContractModelValidation(t *testing.T) {
	// Test Missing Required Field
	testMissingRequiredField(t)

	// Test Too Many Options Contracts
	testTooManyOptionsContracts(t)

	// Test Invalid Options Contract Type
	testInvalidOptionsContractType(t)

	// Test Invalid Options Contract Long Short
	testInvalidOptionsContractLongShort(t)
}

func TestAnalysisEndpoint(t *testing.T) {
	// Happy Path Test
	mockedOptionsContracts := getMockedTestData(FilePathValidOptionsContracts)
	router := initializeTestServer()

	requestPayload, err := json.Marshal(mockedOptionsContracts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/analyze", strings.NewReader(string(requestPayload)))
	router.ServeHTTP(response, request)

	var analysisResponse model_response.AnalysisResponse
	err = json.Unmarshal(response.Body.Bytes(), &analysisResponse)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, 186.41, analysisResponse.MaxProfit)
	assert.Equal(t, -24.35, analysisResponse.MaxLoss)
	assert.Equal(t, 1, len(analysisResponse.BreakEvenPoints))
	assert.Equal(t, 114.73, analysisResponse.BreakEvenPoints[0])
	assert.Equal(t, getAverageNoOfDays(mockedOptionsContracts), len(analysisResponse.XYValues))

	fmt.Println(len(analysisResponse.XYValues))
}
