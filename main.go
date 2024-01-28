package main

import (
	"triumph_intern/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/buy", controllers.Buy)
	app.Get("/sell", controllers.Sell)

	app.Listen(":4000")
}
