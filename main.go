package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/sudheerpal/learn-go/database"
	"github.com/sudheerpal/learn-go/routes"
)

func main() {

	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(logger.New(logger.Config{
		Format: "${time} - ${ip}  - ${status} ${method} ${latency} ${path}\n",
	}))
	app.Use(limiter.New(limiter.Config{
		Max:               5,
		Expiration:        10 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	routes.Setup(app)

	fmt.Println("Hey whats up?")

	app.Listen(":3000")
}
