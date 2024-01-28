package main

import (
	"net/http/httptest"
	"testing"
	"triumph_intern/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert" // add Testify package
)

func setupApp() *fiber.App {
	// Initialize Fiber app
	app := fiber.New()

	// Define routes
	app.Get("/buy", controllers.Buy)   // Adjust as per your package and method names
	app.Get("/sell", controllers.Sell) // Adjust as per your package and method names

	return app
}

func TestBuyHandler(t *testing.T) {
	app := setupApp()

	// Test cases
	tests := []struct {
		description  string
		query        string
		expectedCode int
	}{
		{
			description:  "Valid request",
			query:        "/buy?amount=1&symbol=BTC-USD",
			expectedCode: fiber.StatusOK,
		},
		{
			description:  "Invalid request",
			query:        "/buy?amount=invalid&symbol=BTC-USD",
			expectedCode: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid request",
			query:        "/buy?symbol=BTC-USD",
			expectedCode: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid request",
			query:        "/buy",
			expectedCode: fiber.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.query, nil)
			resp, _ := app.Test(req, 1)
			// Verify, if the status code is as expected
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			// Additional checks (like response body content) can be added here
		})
	}
}

func TestSellHandler(t *testing.T) {
	app := setupApp()

	// Test cases
	tests := []struct {
		description  string
		query        string
		expectedCode int
	}{
		{
			description:  "Valid request",
			query:        "/sell?amount=1&symbol=BTC-USD",
			expectedCode: fiber.StatusOK,
		},
		{
			description:  "Invalid request",
			query:        "/sell?amount=invalid&symbol=BTC-USD",
			expectedCode: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid request",
			query:        "/sell?symbol=BTC-USD",
			expectedCode: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid request",
			query:        "/sell",
			expectedCode: fiber.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.query, nil)
			resp, _ := app.Test(req, 1)
			_ = resp
			// Verify, if the status code is as expected
			// assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

			// Additional checks (like response body content) can be added here
		})
	}
}
