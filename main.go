package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sudheerpal/learn-go/database"
	"github.com/sudheerpal/learn-go/middlewares"
	"github.com/sudheerpal/learn-go/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	middlewares.ApplyMiddleware(app)

	routes.Setup(app)

	app.Listen(":3000")
}
