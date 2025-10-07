package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber/handlers"
	"fiber/middleware"
)

func BookRoutes(app *fiber.App) {
	book := app.Group("/api/books",middleware.JWTProtected())

	book.Post("/",middleware.IsAdmin() ,handlers.CreateBook)
	book.Get("/", handlers.GetAllBooks)
	book.Get("/:id", handlers.GetBookById)
	book.Put("/:id",middleware.IsAdmin(), handlers.UpdateBook)
	book.Delete("/:id",middleware.IsAdmin(), handlers.DeleteBook)
}


