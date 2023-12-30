package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// get sha from -ldflags
var sha string

func main() {
	e := echo.New()

	if os.Getenv("BACKEND_NODE_URL") == "" {
		e.Logger.Error("BACKEND_NODE_URL env is required")
		os.Exit(1)
	}

	// Routes
	e.GET("/healthz", healthcheck)

	e.GET("/version", version)

	e.POST("/call-node-backend/add", func(c echo.Context) error {
		return makeRequest(c, "add")
	})

	e.POST("/call-node-backend/subtract", func(c echo.Context) error {
		return makeRequest(c, "subtract")
	})

	e.POST("/call-node-backend/multiply", func(c echo.Context) error {
		return makeRequest(c, "multiply")
	})

	e.POST("/call-node-backend/divide", func(c echo.Context) error {
		return makeRequest(c, "divide")
	})

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}

func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "The go service is healthy")
}

func version(c echo.Context) error {
	version := fmt.Sprintf("app_sha: %s", sha)
	return c.String(http.StatusOK, version)
}

// Function to make HTTP requests to the Node.js backend
func makeRequest(c echo.Context, operation string) error {
	// Parse the request body to get the operands
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "Error parsing request body")
	}

	// Make a request to the Node.js backend to perform the calculation
	nodeBackendURL := fmt.Sprintf("%s/api/%s", os.Getenv("BACKEND_NODE_URL"), operation)

	nodeResponseBody, err := makeRequestHelper(http.MethodPost, nodeBackendURL, requestBody)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error calling Node.js backend")
	}

	return c.JSONBlob(http.StatusOK, nodeResponseBody)
}

// Function to make HTTP requests
func makeRequestHelper(method, url string, requestBody map[string]interface{}) ([]byte, error) {
	requestBodyBytes, _ := json.Marshal(requestBody)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
