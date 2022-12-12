package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ApplyMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:               5,
		Expiration:        10 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
}
