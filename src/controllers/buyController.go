package controllers

import (
	"triumph_intern/services"

	"github.com/gofiber/fiber/v2"
)

func Buy(c *fiber.Ctx) error {
	// Extract query parameters
	amount := c.Query("amount")
	symbol := c.Query("symbol")

	requestValid := services.ValidateRequest(amount, symbol)
	if requestValid != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": requestValid.Error(),
		})
	}
	// Call service
	order, err := services.ExecuteOrder(amount, symbol, "buy")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"order":  order,
	})
}
