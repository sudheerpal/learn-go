package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudheerpal/learn-go/controllers"
	"github.com/sudheerpal/learn-go/middlewares"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	// test the system
	api.Get("/", controllers.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", controllers.Logout)

	//User
	user := api.Group("/user")
	user.Post("/register", controllers.Register)
	user.Get("/get", middlewares.Protected, controllers.User)
	user.Get("/all", controllers.AllUsers)

	//Lead
	lead := api.Group("/lead", middlewares.Protected)
	lead.Get("/all", controllers.AllLeads)

}
