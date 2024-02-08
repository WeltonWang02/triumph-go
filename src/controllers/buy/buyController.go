package buy

import (
	"triumph_intern/enum"
	"triumph_intern/services"

	"github.com/gofiber/fiber/v2"
)

func Buy(c *fiber.Ctx) error {
	// Extract query parameters
	amount := c.Query("amount")
	symbol := c.Query("symbol")

	if requestValid := services.ValidateRequest(amount, symbol); requestValid != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  enum.Error,
			"message": requestValid.Error(),
		})
	}
	// Call service
	order, err := services.ExecuteOrder(amount, symbol, "buy")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  enum.Error,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": enum.Success,
		"order":  order,
	})
}
