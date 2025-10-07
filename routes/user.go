package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber/handlers"
)

func UserRoutes(app *fiber.App) {
	app.Post("/api/register", handlers.RegisterUser)
	app.Post("/api/login", handlers.LoginUser)
}


