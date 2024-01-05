package routes

import (
	"github.com/gofiber/fiber/v2"
	"server/pkg/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.HandleRegister)
	app.Post("/api/login", controllers.HandleLogin)
	app.Get("/health", controllers.HandleHealth)
	app.Post("/contact-us", controllers.HandleContactUs)
	// app.Get("/login", controllers.HandleLoginView)
	// app.Get("/register", controllers.HandleRegisterView)
}
