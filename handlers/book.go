package handlers

import (
	"context"
	"fiber/config"
	"fiber/models"
	"fiber/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	validationErrors := utils.ValidateStruct(book)
	if validationErrors != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": validationErrors,
		})
	}


	book.ID = primitive.NewObjectID()

	collection := config.DB.Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to insert book"})
	}

	return c.Status(201).JSON(book)
	// return c.Status(201).JSON(fiber.Map{"book": book, "data": data})
}

func GetAllBooks(c *fiber.Ctx) error {
	collection := config.DB.Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, primitive.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch books"})
	}
	defer cursor.Close(ctx)

	var books []models.Book
	if err = cursor.All(ctx, &books); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to parse books"})
	}

	return c.Status(200).JSON(books)
}

func GetBookById(c *fiber.Ctx) error{

	if c.Params("id") == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Missing book ID"})
	}

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	collection := config.DB.Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book models.Book
	err = collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&book)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.Status(200).JSON(book)

}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updatedBook models.Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// validationErrors := utils.ValidateStruct(updatedBook)
	// if validationErrors != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"error":   "Validation failed",
	// 		"details": validationErrors,
	// 	})
	// }


	collection := config.DB.Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":     updatedBook.Title,
			"author":    updatedBook.Author,
			"publisher": updatedBook.Publisher,
			"year":      updatedBook.Year,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update book"})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.JSON(fiber.Map{"message": "Book updated successfully"})
}


func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	collection := config.DB.Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.JSON(fiber.Map{"message": "Book deleted successfully"})
}