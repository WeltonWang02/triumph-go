package main

import (
	"triumph_intern/controllers/buy"
	"triumph_intern/controllers/sell"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/buy", buy.Buy)
	app.Get("/sell", sell.Sell)

	app.Listen(":4000")
}
