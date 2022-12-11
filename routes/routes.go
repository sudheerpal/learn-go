package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sudheerpal/learn-go/controllers"
)

func Setup(app *fiber.App) {

	api := app.Group("/api", logger.New())

	api.Get("/", controllers.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", controllers.Logout)

	//User
	user := api.Group("/user")
	user.Post("/register", controllers.Register)
	user.Get("/get", controllers.User)
	user.Get("/all", controllers.AllUsers)

	//Lead
	// lead := api.Group("/lead")

}
