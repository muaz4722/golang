package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber/handlers"
)

func BookRoutes(app *fiber.App) {
	book := app.Group("/api/books")

	book.Post("/", handlers.CreateBook)
	book.Get("/", handlers.GetAllBooks)
	book.Get("/:id", handlers.GetBookById)
	book.Put("/:id", handlers.UpdateBook)
	book.Delete("/:id", handlers.DeleteBook)
}
