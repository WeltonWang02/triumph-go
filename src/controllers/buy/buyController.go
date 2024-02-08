package buy

import (
	"triumph_intern/enum"
	"triumph_intern/models"
	"triumph_intern/services"

	"github.com/gofiber/fiber/v2"
)

func Buy(c *fiber.Ctx) error {
	// Extract query parameters
	orderRequest := models.OrderRequest{
		Amount: c.Query("amount"),
		Symbol: c.Query("symbol"),
	}

	if err := orderRequest.ValidateRequest(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  enum.Error,
			"message": err.Error(),
		})
	}
	// Call service
	order, err := services.ExecuteOrder(orderRequest, "buy")
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
