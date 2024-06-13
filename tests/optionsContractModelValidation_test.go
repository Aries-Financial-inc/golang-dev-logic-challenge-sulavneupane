package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-sulavneupane/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func callAnalyzeEndpointWithErrorResponse(t *testing.T, payloadBytes []byte) (*httptest.ResponseRecorder, errors.ErrorResponse) {
	router := initializeTestServer()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/analyze", strings.NewReader(string(payloadBytes)))
	router.ServeHTTP(response, request)

	var errorResponse errors.ErrorResponse
	err := json.Unmarshal(response.Body.Bytes(), &errorResponse)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return response, errorResponse
}

func testMissingRequiredField(t *testing.T) {
	// Test Missing Required Field
	mockedOptionsContracts := getMockedTestData(FilePathMissingRequiredField)
	requestPayload, err := json.Marshal(mockedOptionsContracts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	router := initializeTestServer()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/analyze", strings.NewReader(string(requestPayload)))
	router.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code)
}

func testTooManyOptionsContracts(t *testing.T) {
	// Test Too Many Options Contracts
	mockedOptionsContracts := getMockedTestData(FilePathTooManyOptionsContracts)
	requestPayload, err := json.Marshal(mockedOptionsContracts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	response, errorResponse := callAnalyzeEndpointWithErrorResponse(t, requestPayload)
	assert.Equal(t, 400, response.Code)
	assert.Equal(t, errors.ErrorCodeTooManyOptionsContracts, errorResponse.ErrorCode)
	assert.Equal(t, "too many option contracts", errorResponse.Message)
}

func testInvalidOptionsContractType(t *testing.T) {
	// Test Invalid Options Contract Type
	mockedOptionsContracts := getMockedTestData(FilePathInvalidOptionsContractType)
	requestPayload, err := json.Marshal(mockedOptionsContracts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	response, errorResponse := callAnalyzeEndpointWithErrorResponse(t, requestPayload)
	assert.Equal(t, 400, response.Code)
	assert.Equal(t, errors.ErrorCodeInvalidOptionsContractType, errorResponse.ErrorCode)
	assert.Equal(t, "invalid contract option type caller", errorResponse.Message)
}

func testInvalidOptionsContractLongShort(t *testing.T) {
	// Test Invalid Options Contract Long Short
	mockedOptionsContracts := getMockedTestData(FilePathInvalidOptionsContractLongShort)
	requestPayload, err := json.Marshal(mockedOptionsContracts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	response, errorResponse := callAnalyzeEndpointWithErrorResponse(t, requestPayload)
	assert.Equal(t, 400, response.Code)
	assert.Equal(t, errors.ErrorCodeInvalidOptionsContractLongShort, errorResponse.ErrorCode)
	assert.Equal(t, "invalid contract option long/short: longer", errorResponse.Message)
}
