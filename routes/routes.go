package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudheerpal/learn-go/controllers"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)
	app.Get("/api/users", controllers.AllUsers)

}
