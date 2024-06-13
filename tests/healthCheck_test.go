package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func healthCheckHappyPathTest(t *testing.T) {
	// Happy Path Test
	router := initializeTestServer()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/health-check", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "OK", response.Body.String())
}
