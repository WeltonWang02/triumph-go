package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TimerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		log.Println("Request took ", stop.Sub(start).String())

		return nil
	}
}
