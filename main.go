package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/sudheerpal/learn-go/database"
	"github.com/sudheerpal/learn-go/routes"
)

func main() {

	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	fmt.Println("Hey whats up?")

	app.Listen(":3000")
}
