package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
	e := echo.New()

	e.GET("/healthz", healthcheck)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "The go service is healthy", rec.Body.String())
}

func TestAddEndpoint(t *testing.T) {
	operation_type := "add"
	endpoint := fmt.Sprintf("/call-node-backend/%s", operation_type)
	requestBody := map[string]interface{}{
		"num1": 3,
		"num2": 5,
	}

	e := echo.New()

	e.POST(endpoint, func(c echo.Context) error {
		return makeRequest(c, operation_type)
	})

	// Prepare the request
	reqBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var result map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &result)
	assert.NoError(t, err)

	// Importantly, use assert.EqualValues to compare float64 values
	assert.EqualValues(t, int64(8), result["result"])
}

func TestSubtractEndpoint(t *testing.T) {
	operation_type := "subtract"
	endpoint := fmt.Sprintf("/call-node-backend/%s", operation_type)
	requestBody := map[string]interface{}{
		"num1": 3,
		"num2": 5,
	}

	e := echo.New()

	e.POST(endpoint, func(c echo.Context) error {
		return makeRequest(c, operation_type)
	})

	// Prepare the request
	reqBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var result map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &result)
	assert.NoError(t, err)

	// Importantly, use assert.EqualValues to compare float64 values
	assert.EqualValues(t, int64(-2), result["result"])
}

func TestMultiplyEndpoint(t *testing.T) {
	operation_type := "multiply"
	endpoint := fmt.Sprintf("/call-node-backend/%s", operation_type)
	requestBody := map[string]interface{}{
		"num1": 3,
		"num2": 5,
	}

	// Create an instance of echo.Echo
	e := echo.New()

	e.POST(endpoint, func(c echo.Context) error {
		return makeRequest(c, operation_type)
	})

	// Prepare the request
	reqBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var result map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &result)
	assert.NoError(t, err)

	// Importantly, use assert.EqualValues to compare float64 values
	assert.EqualValues(t, int64(15), result["result"])
}

func TestDivideEndpoint(t *testing.T) {
	operation_type := "divide"
	endpoint := fmt.Sprintf("/call-node-backend/%s", operation_type)
	requestBody := map[string]interface{}{
		"num1": 100,
		"num2": 20,
	}

	// Create an instance of echo.Echo
	e := echo.New()

	e.POST(endpoint, func(c echo.Context) error {
		return makeRequest(c, operation_type)
	})

	// Prepare the request
	reqBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var result map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &result)
	assert.NoError(t, err)

	// Importantly, use assert.EqualValues to compare float64 values
	assert.EqualValues(t, int64(5), result["result"])
}
